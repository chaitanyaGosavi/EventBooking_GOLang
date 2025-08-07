# Golang API with Gin, SQLite, JWT & Docker

This is a simple RESTful API built using **Go (Golang)** and the **Gin** framework. It supports user registration, authentication with JWT, event creation and listing, and uses **SQLite** as the database.

## Features

- RESTful API using Gin
- SQLite for lightweight persistent storage
- JWT-based authentication
- Password hashing with bcrypt

## Technologies Used

- Go (Golang)
- Gin Web Framework
- SQLite3 Database
- JWT (golang-jwt)
- Bcrypt for password hashing

## Project Structure

.
├── db/ # Database setup and queries
├── middlewares/ # Auth, logging, etc.
├── models/ # Data models (User, Event)
├── routes/ # API route handlers
├── utils/ # Utility functions like token generation
├── main.go # Application entry point
├── go.mod
├── go.sum
├── Dockerfile
└── README.md
