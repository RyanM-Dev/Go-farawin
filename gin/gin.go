package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//You can specify the port on the command line

var port = flag.Int("port", 8080, "Port to run the HTTP server")

//Greeting Method

func Hello(c *gin.Context) {
	c.String(200, "Hello Ryan and Mr.Ghofrani")
}
func apisPrint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Apis": []string{
			"GET   /hello",
			"POST   /print",
			"DELETE  /stop",
		},
	})
}

//Auth Method

func auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "test" {
		c.AbortWithStatus(http.StatusForbidden)
		c.String(http.StatusForbidden, "Authentication Required")
		log.Print(http.StatusForbidden, "  Authentication Required")
		return
	}
	c.Next()
}

//Query and Body print Method,Body has priority

func printMessageHandler(c *gin.Context) {

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")

		fmt.Printf("error : %s", err)
		return
	} else if string(body) != "" {
		var jsonData interface{}
		err = json.Unmarshal(body, &jsonData)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error parsing JSON",
			})
			return
		} else if err == nil {
			c.JSON(http.StatusOK, jsonData)
		}

	} else {

		urlParam := c.Query("msg")

		if urlParam != "" {
			c.String(http.StatusOK, "Your Message is: %s", urlParam)
		} else {
			c.String(http.StatusGone, "no message received ")
		}

	}
}

func main() {

	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	router := gin.Default()
	router.Use(auth)
	router.GET("/hello", Hello)
	router.GET("/", apisPrint)
	router.POST("/print", printMessageHandler)

	router.DELETE("/stop", func(c *gin.Context) {
		os.Exit(0)
	})
	router.Run(addr)

}
