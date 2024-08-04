package controllers

import (
	"net/http"
	"quiz-3/database"
	"quiz-3/repository"
	"quiz-3/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	// authService services.AuthService
	authRepository repository.UserRepository
}

// func NewAuthController(authService services.AuthService) *AuthController {
func NewAuthController(ar repository.UserRepository) *AuthController {
	// return &AuthController{authService: authService}
	return &AuthController{authRepository: ar}
}

func Login(ctx *gin.Context) {
	// ac := controllers.NewAuthController(repository.NewUserRepository(database.DbConnection))

	userRepository := repository.NewUserRepository(database.DbConnection)
	authService := services.NewAuthService(userRepository)
	// authController := NewAuthController(authService.)

	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username dan Password tidak boleh kosong"})
		return
	}

	// user, err := authController.authService.Authenticate(username, password)
	user, err := authService.Authenticate(username, password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing credentials"})
		return
	}
	// Login successful, return user data or token
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
