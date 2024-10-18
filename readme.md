### Structure Project

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
|   |   |-- bcrypt.go          # Bcrypt implementation
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
