package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	fmt.Println("Everything OK")
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func usefullLink(c *gin.Context) {
	fmt.Println("Jaja, works")
	c.Writer.Header().Add("all", "fine")
}
