package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"quiz-3/controllers"
	"quiz-3/database"
	"quiz-3/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("error loading env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	database.DBMigrate(DB)

	// userRepository := repository.NewUserRepository(DB)
	// authService := services.NewAuthService(userRepository)
	// authController := controllers.NewAuthController(authService)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.POST("/login", controllers.Login)

	// Category

	fmt.Println(router)
	router.Run(":8080")
}

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func VerifyPassword(password, hashedPassword string) bool {
	h := sha256.New()
	h.Write([]byte(password))
	return hashedPassword == base64.StdEncoding.EncodeToString(h.Sum(nil))
}
