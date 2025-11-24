package services

import (
	"errors"
	"os"
	"time"

	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

type AuthService interface {
	Register(user *models.User) error
	Login(email, password string) (string, error)
	GoogleLogin(userInfo *GoogleUserInfo) (string, *models.User, error)
}

type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(user *models.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.repo.Create(user)
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := s.generateJWT(user)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}

func (s *authService) GoogleLogin(userInfo *GoogleUserInfo) (string, *models.User, error) {
	user, err := s.repo.FindByGoogleID(userInfo.ID)
	if err == nil {
		token, err := s.generateJWT(user)
		if err != nil {
			return "", nil, errors.New("failed to generate token")
		}
		return token, user, nil
	}

	user, err = s.repo.FindByEmail(userInfo.Email)
	if err == nil {
		user.GoogleID = &userInfo.ID
		user.Provider = "google"
		user.ProfileImage = &userInfo.Picture

		if err := s.repo.Update(user); err != nil {
			return "", nil, errors.New("failed to link Google account to existing account")
		}

		token, err := s.generateJWT(user)
		if err != nil {
			return "", nil, errors.New("failed to generate token")
		}
		return token, user, nil
	}

	newUser := &models.User{
		Email:        userInfo.Email,
		GoogleID:     &userInfo.ID,
		Provider:     "google",
		Username:     &userInfo.Name,
		ProfileImage: &userInfo.Picture,
		Password:     "",
	}

	if err := s.repo.Create(newUser); err != nil {
		return "", nil, errors.New("failed to create user from Google account")
	}

	token, err := s.generateJWT(newUser)
	if err != nil {
		return "", nil, errors.New("failed to generate token")
	}
	return token, newUser, nil
}

func (s *authService) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-secret-key-change-this"
	}

	return token.SignedString([]byte(secret))
}
