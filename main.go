package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func constantLoop(seconds int) {
	for {
		time.Sleep(time.Second * time.Duration(seconds))
		resp, err := http.Get("http://localhost:8080/api/v1/usefull_link")
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(time.Now().Local().Format("00:00:00") + " -> Status: " + resp.Status)
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

var startTime time.Time

func printUptime(c *gin.Context) {
	elapsedTime := time.Since(startTime)
	fmt.Println(elapsedTime)
}

func main() {
	router := gin.New()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./*.html")

	startTime = time.Now()

	go constantLoop(30)

	router.GET("/", rootHandler)
	router.GET("/api/v1/usefull_link", usefullLink)
	router.GET("/api/v1/uptime", printUptime)

	if err := router.Run(":8080"); err != nil {
		log.Panicf("[-] Error: %s", err)
	}
}
