package middleware

import (
	"net/http"
	"quiz-3/database"
	"quiz-3/repository"
	"quiz-3/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRepository := repository.NewUserRepository(database.DbConnection)
		authService := services.NewAuthService(userRepository)

		username, password, ok := ctx.Request.BasicAuth()

		if username == "" || password == "" || !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing credentials"})
			ctx.Abort()
			return
		}

		// Call the auth controller to authenticate
		user, err := authService.Authenticate(username, password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing credentials"})
			ctx.Abort()
			return
		}

		// Set the authenticated user in the context
		ctx.Set("user", user)
		ctx.Next()
	}
}
