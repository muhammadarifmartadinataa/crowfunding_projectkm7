package handler

import (
	"crowfundig/gemini"
	"crowfundig/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeminiHandler struct {
	service gemini.GeminiService
}

func NewGeminiHandler(service gemini.GeminiService) *GeminiHandler {
	return &GeminiHandler{service}
}

func (h *GeminiHandler) SaveGeminiResponse(c *gin.Context) {
	// Panggil API Gemini
	response, err := utils.CallGeminiAPI()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get response from Gemini API"})
		return
	}

	// Simpan data ke database
	savedResponse, err := h.service.SaveResponse(response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    savedResponse,
	})
}

func (h *GeminiHandler) GetGeminiResponses(c *gin.Context) {
	geminiResponses, err := h.service.GetAllResponses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    geminiResponses,
	})
}
