package api

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var port = flag.Int("port", 8080, "Port to run the HTTP server")

func RunApi() {

	flag.Parse()
	addr := fmt.Sprintf(":%d", *port)
	router := gin.Default()
	router.Use(Auth)
	router.GET("/hello", Hello)
	router.GET("/", ApisPrint)
	router.POST("/print", PrintMessageHandler)

	router.DELETE("/stop", func(c *gin.Context) {
		os.Exit(0)
	})
	router.Run(addr)

}
