package main

import (
	"chathenon/db"
	"chathenon/handler"
	"chathenon/middleware"
	"chathenon/repository"
	"chathenon/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	r.Use(middleware.ErrorMiddleware())

	db := db.InitDB()
	if db == nil {
		panic("Failed to connect to the database")
	}

	// Repository
	ur := repository.NewUserRepo(db)

	// Use Case
	uuc := usecase.NewUserUseCase(ur)

	// Handler
	uh := handler.NewUserHandler(uuc)

	r.POST("/users/register", uh.RegisterHandler)
	r.POST("/users/login", uh.LoginHandler)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":8080")
}
