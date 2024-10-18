package routes

import (
	"myapp/internal/user"
	"myapp/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userHandler *user.Handler) {
	// Public routes
	router.POST("/register", userHandler.Register)
	router.POST("/login", userHandler.Login)

	// Protected routes (JWT Authentication required)
	protected := router.Group("/users")
	protected.Use(jwt.JWTAuthMiddleware())
	protected.GET("/", userHandler.GetUsers)
}
