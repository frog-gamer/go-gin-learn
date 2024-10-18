# MyApp

## Overview

**MyApp** is a user management API built using **Go** with **Gin** for routing, **JWT** for authentication, **AES-256** encryption for securing sensitive data, and **bcrypt** for password hashing. It follows **Clean Architecture** principles to ensure a scalable and maintainable codebase. The application also supports features such as pagination and leverages **PostgreSQL** for data persistence.

## Project Structure

```
/myapp
|-- /cmd
|   |-- main.go                 # Main entry point of the Go application
|-- /config
|   |-- config.go               # Config file to manage environment variables
|-- /internal
|   |-- /user
|       |-- handler.go          # HTTP handler for user routes
|       |-- repository.go       # Interface for user repository
|       |-- repository_impl.go  # PostgreSQL repository implementation
|       |-- usecase.go          # Interface for user use case
|       |-- usecase_impl.go     # Implementation for user use cases
|       |-- entity.go           # User entity definition
|-- /pkg
|   |-- /crypto
|   |   |-- aes.go              # AES encryption and decryption
|   |   |-- bcrypt.go           # Bcrypt implementation
|   |-- /jwt
|   |   |-- jwt.go              # JWT generation and validation
|   |-- /response
|   |   |-- response.go         # Standardized JSON response utility
|-- /routes
|   |-- user_routes.go          # Route definitions for user-related routes
|-- .env                        # Environment variables
|-- Dockerfile                  # Dockerfile to containerize the Go application
|-- docker-compose.yml          # Docker Compose setup for the app and database
|-- go.mod                      # Go modules file with updated dependencies
|-- go.sum                      # Go modules checksum file
```

## Features

- **Clean Architecture**: Separation of concerns between handlers, use cases, repositories, and entities.
- **Gin**: Lightweight web framework for efficient HTTP routing.
- **JWT Authentication**: Token-based authentication system.
- **AES-256 Encryption**: Secure encryption for handling sensitive data.
- **Bcrypt**: Secure password hashing for user credentials.
- **Pagination**: Efficient data pagination in the repository layer.
- **PostgreSQL**: Relational database used for persistent storage.
- **Dockerized**: Ready for containerization and easy deployment using Docker and Docker Compose.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/frog-gamer/go-gin-learn.git
   cd myapp
   ```

2. Set up environment variables in `.env`:

   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=youruser
   DB_PASSWORD=yourpassword
   DB_NAME=myappdb
   JWT_SECRET=yourjwtsecret
   AES_SECRET=youraessecret
   ```

3. Install Go dependencies:

   ```bash
   go mod download
   ```

4. Run the application:

   ```bash
   go run cmd/main.go
   ```

## Docker Setup

1. Build and start the app using Docker Compose:

   ```bash
   docker-compose up --build
   ```

2. This will start the Go application and PostgreSQL database in containers. The API will be accessible at `http://localhost:8080`.

## Endpoints

| Method | Endpoint           | Description                     |
|--------|--------------------|---------------------------------|
| POST   | `/users/register`   | Register a new user             |
| POST   | `/users/login`      | Login and receive JWT token     |
| GET    | `/users`            | Get list of users (pagination)  |
| GET    | `/users/:id`        | Get user by ID                  |
| PUT    | `/users/:id`        | Update user by ID               |
| DELETE | `/users/:id`        | Delete user by ID               |

## Testing

To run tests, use the following command:

```bash
go test ./...
```

## Security

- **JWT** is used for stateless user authentication.
- **AES-256** ensures sensitive data is encrypted.
- **Bcrypt** hashes user passwords before storing them in the database.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
