# GORM Learning Project

This project is a simple Go application designed for learning purposes. It demonstrates how to use the GORM library to interact with a MySQL database. The project includes examples of defining models, relationships, and performing basic CRUD operations.

## Features

- Models with relationships:
	- `Product` with `Category` (many-to-many)
	- `Product` with `SerialNumber` (one-to-one)
- Database migrations using `AutoMigrate`
- Preloading related data for queries

## Prerequisites

- Go installed on your machine
- MySQL database running locally
- GORM library installed

## Setup

1. Clone the repository.
2. Update the `dsn` variable in `main.go` with your MySQL credentials.
3. Run the application:
	 ```bash
	 go run main.go
	 ```

## Purpose

This project is solely for educational purposes to understand how to work with GORM in Go.