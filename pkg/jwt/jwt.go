package jwt

import (
	"log"
	"myapp/config"
	"myapp/pkg/response"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Claims struct to embed custom JWT claims
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token for a user with a given user ID
func GenerateToken(userID int) (string, error) {
	// Set custom claims
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)), // Token valid for 72 hours
		},
	}

	// Create a new JWT token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(config.GetJWTSecretKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// JWTAuthMiddleware is a middleware for validating JWT tokens
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		//log.Println("Authorization header:", tokenString) // Log the Authorization header

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:] // Remove "Bearer " prefix
		} else {
			response.ErrorResponse(c, "401", "Missing or invalid token")
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJWTSecretKey()), nil
		})

		if err != nil || !token.Valid {
			log.Println("Invalid token:", err) // Log the token error for debugging
			response.ErrorResponse(c, "401", "Invalid token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Claims)
		if ok && token.Valid {
			log.Println("User ID from token:", claims.UserID) // Log the extracted user ID
			c.Set("user_id", claims.UserID)
		} else {
			response.ErrorResponse(c, "401", "Invalid token claims")
			c.Abort()
			return
		}

		c.Next()
	}
}

// ExtractUserID extracts the user ID from the JWT claims stored in the context
func ExtractUserID(c *gin.Context) (int, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}

	userIDInt, ok := userID.(int)
	return userIDInt, ok
}
