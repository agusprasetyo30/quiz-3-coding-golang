package main

import (
	"database/sql"
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
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
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

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.POST("/api/login", controllers.Login)

	// Category
	router.GET("/api/categories", controllers.GetAllCategory)
	router.GET("/api/categories/:id/books", controllers.GeBookByCategory)
	router.POST("/api/categories", controllers.InsertCategory)
	router.GET("/api/categories/:id", controllers.GetCategory)
	router.PUT("/api/categories/:id", controllers.UpdateCategory)
	router.DELETE("/api/categories/:id", controllers.DeleteCategory)

	// Book
	router.GET("/api/books", controllers.GetAllBook)
	router.POST("/api/books", controllers.InsertBook)
	router.GET("/api/books/:id", controllers.GetBook)
	router.PUT("/api/books/:id", controllers.UpdateBook)
	router.DELETE("/api/books/:id", controllers.DeleteBook)

	fmt.Println(router)
	router.Run(":8080")
}
