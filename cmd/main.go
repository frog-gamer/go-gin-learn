package main

import (
	"database/sql"
	"log"
	"myapp/config"
	"myapp/internal/user"
	"myapp/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// Health check route
	r.GET("/health", HealthCheck)

	// Initialize the database
	db := initDB()
	defer db.Close()

	// Initialize repository, use case, and handler for user
	userRepo := user.NewPGRepository(db)
	userUseCase := user.NewUseCase(userRepo)
	userHandler := user.NewHandler(userUseCase)

	// Register user routes
	routes.UserRoutes(r, userHandler)

	// Start the server
	if err := r.Run(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

// HealthCheck is a handler function for the health check endpoint
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"responseCode":    "200",
		"responseMessage": "Service is up",
	})
}

func initDB() *sql.DB {
	db, err := sql.Open("postgres", config.GetDBConnectionString())
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database is not responding: ", err)
	}

	return db
}
