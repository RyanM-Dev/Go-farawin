package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Greeting Method

func Hello(c *gin.Context) {
	c.String(200, "Hello Ryan and Mr.Ghofrani")
}
func ApisPrint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Apis": []string{
			"GET   /hello",
			"POST   /print",
			"DELETE  /stop",
		},
	})
}

//Auth Method

func Auth(c *gin.Context) {
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

func PrintMessageHandler(c *gin.Context) {

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
