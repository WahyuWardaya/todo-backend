package main

import (
	"log"
	"time"

	"todo-backend/database"
	"todo-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	// Konfigurasi CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Ganti dengan URL frontend Anda
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))

	// Rute API
	router.GET("/todos", handlers.GetTodos)
	router.POST("/todos", handlers.CreateTodo)
	router.PUT("/todos/:id", handlers.UpdateTodo)
	router.DELETE("/todos/:id", handlers.DeleteTodo)

	log.Println("Server berjalan di http://localhost:8080")
	router.Run(":8080")
}