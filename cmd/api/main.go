package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaeloduarte/url-shortener-go/internal/models"
	"github.com/rafaeloduarte/url-shortener-go/internal/routes"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
