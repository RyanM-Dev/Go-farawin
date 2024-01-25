package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.String(200, "Hello Ryan and Mr.Ghofrani")
}
func apisPrint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Apis": []string{
			"GET   /hello",
			"POST   /print",
			"DELETE  /stop",
			"POST  /print",
		},
	})
}
func auth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader != "test" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Next()
}
func printMessageHandler(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")

		fmt.Printf("error : %s", err)
		return
	} else if string(body) != "" {
		c.JSON(http.StatusOK, gin.H{
			"body": string(body),
		})

	} else {

		urlParam := c.Query("msg")

		if urlParam != "" {
			c.String(http.StatusOK, "Your Message is: %s", urlParam)
		} else {
			c.String(http.StatusGone, "no message received ")
		}
		var jsonBody map[string]interface{}
		err = c.ShouldBindJSON(&jsonBody)
		if jsonBody != nil {
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, jsonBody)
		}

	}
}

var port = flag.Int("port", 8080, "Port to run the HTTP server")

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
