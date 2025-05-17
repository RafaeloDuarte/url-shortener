package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rafaeloduarte/url-shortener-go/internal/models"
)

const shortCodeLength = 6

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generateShortCode() string {
	rand.Seed(time.Now().UnixNano())

	code := make([]rune, shortCodeLength)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}

	return string(code)
}

func ShortenURL(c *gin.Context) {
	var request struct {
		OriginalURL string `json:"original_url" binding:"required,url"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL inválida"})
		return
	}

	shortCode := generateShortCode()

	url := models.URL{
		OriginalURL: request.OriginalURL,
		ShortCode:   shortCode,
		Clicks:      0,
	}

	if err := models.DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_url": c.Request.Host + "/" + shortCode,
	})
}

func RedirectToOriginalURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	var url models.URL
	if err := models.DB.First(&url, "short_code = ?", shortCode).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Código curto não encontrado"})
		return
	}

	models.DB.Model(&url).Update("clicks", url.Clicks+1)

	c.Redirect(http.StatusFound, url.OriginalURL)
}
