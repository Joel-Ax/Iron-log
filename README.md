# Go Fiber Postgres Workout Tracker

A RESTful API built with Go, Fiber, and PostgreSQL for managing workout tracking and exercise data.

## Purpose

The purpose of this application is to provide a backend for a workout progress app. It enables users to track their fitness journey by logging workouts, managing exercise libraries, recording sets and reps, and monitoring progress over time. The API is designed to be consumed by web or mobile frontends, providing a robust foundation for comprehensive fitness tracking applications.

## Overview

This application provides a backend service for tracking workouts, exercises, and exercise sets. It follows a clean architecture pattern with separated layers for controllers, services, repositories, and models, making it maintainable and scalable.

## Tech Stack

### Fiber Framework

Fiber is an Express-inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. This project uses Fiber because it:

Provides extremely fast HTTP routing and performance
Offers an intuitive, Express-like API that's easy to learn
Has minimal memory allocation and overhead
Includes built-in middleware support
Simplifies JSON parsing and response handling

### GORM

GORM is a developer-friendly ORM library for Go. This project uses GORM to:

Simplify database operations with idiomatic Go code
Handle database migrations automatically
Provide type-safe database queries
Support PostgreSQL with proper connection pooling
Manage relationships between entities (users, workouts, exercises)

## Project Structure:

```.
├── controllers/ # HTTP request handlers
├── services/ # Business logic layer
├── repositories/ # Data access layer
├── models/ # Database models and schemas
├── routes/ # API route definitions
└── storage/ # Database configuration and connection
```

## Architecture Layers

Controllers handle HTTP requests and responses, delegating business logic to services.
Services contain business logic and validation rules, acting as an intermediary between controllers and repositories.
Repositories manage database operations and queries, providing a clean abstraction over GORM.
Models define the database schema using GORM tags for automatic migration and validation.
