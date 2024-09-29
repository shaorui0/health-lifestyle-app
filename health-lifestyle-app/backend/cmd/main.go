package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "os"
    "github.com/yourusername/health-lifestyle-app/backend/internal/handlers"
)

func main() {
    // Load environment variables
    err := godotenv.Load("../.env")
    if err != nil {
        log.Println("No .env file found")
    }

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to the Health & Lifestyle App Backend!",
        })
    })

    // AI Route
    r.POST("/ai/generate", handlers.GetAIResponseHandler)

    // TODO: Register routes for different services

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    r.Run(":" + port)
}
