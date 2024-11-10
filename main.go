package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func constantLoop(seconds int) { // It's supposed that this will make render non to close my service
	for {
		time.Sleep(time.Second * time.Duration(seconds))
		resp, err := http.Get("http://localhost:8080/api/usefull_link")
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(time.Now().Local().Format("15:04:05") + " -> Status: " + resp.Status)
	}
}

func rootHandler(c *gin.Context) {
	fmt.Println("Everything OK")
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func usefullLink(c *gin.Context) {
	fmt.Println("Jaja, works")
	c.Writer.Header().Add("all", "fine")
}

func main() {
	router := gin.New()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./resources/*.html")

	go constantLoop(30)

	router.GET("/", rootHandler)
	router.GET("/api/usefull_link", usefullLink)

	if err := router.Run(":8080"); err != nil {
		log.Panicf("[-] Error: %s", err)
	}
}
