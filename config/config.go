package config

import (
	"os"
)

func GetDBConnectionString() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	return "user=" + user + " password=" + password + " dbname=" + dbname + " host=" + host + " port=" + port + " sslmode=disable"
}

func GetJWTSecretKey() string {
	return os.Getenv("JWT_SECRET") // Ensure this is a 32-character key for AES-256 encryption
}
