package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaeloduarte/url-shortener-go/internal/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Bem-vindo ao encurtador de URL em Go!"})
	})

	r.POST("/shorten", controllers.ShortenURL)
	r.GET("/:shortCode", controllers.RedirectToOriginalURL)
}
