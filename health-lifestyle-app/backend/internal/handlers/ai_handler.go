package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/yourusername/health-lifestyle-app/backend/internal/ai"
)

type AIRequest struct {
    Prompt string `json:"prompt" binding:"required"`
}

type AIResponse struct {
    Response string `json:"response"`
}

func GetAIResponseHandler(c *gin.Context) {
    var req AIRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    aiResponse, err := ai.GetAIResponse(req.Prompt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get AI response"})
        return
    }

    c.JSON(http.StatusOK, AIResponse{Response: aiResponse})
}
