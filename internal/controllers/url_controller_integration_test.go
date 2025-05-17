package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafaeloduarte/url-shortener-go/internal/models"
)

func TestShortenURL(t *testing.T) {
	models.ConnectDatabase() // necessário para testar com DB real (use um DB de teste!)

	router := gin.Default()
	router.POST("/shorten", ShortenURL)

	body := map[string]string{"original_url": "https://google.com"}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Status esperado 200, mas veio %d", w.Code)
	}
}
