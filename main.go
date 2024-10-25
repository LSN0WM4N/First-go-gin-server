package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusOK, "./resources/index.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
