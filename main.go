package main

import (
	"log"
	"os"

	"github.com/Joel-Ax/go-fiber-postgres/controllers"
	"github.com/Joel-Ax/go-fiber-postgres/models"
	"github.com/Joel-Ax/go-fiber-postgres/repositories"
	"github.com/Joel-Ax/go-fiber-postgres/routes"
	"github.com/Joel-Ax/go-fiber-postgres/services"
	"github.com/Joel-Ax/go-fiber-postgres/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Could not load the database")
	}

	err = models.MigrateUsers(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	app := fiber.New()
	routes.SetupRoutes(app, userController)
	app.Listen(":8080")

}
