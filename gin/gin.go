package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.String(200, "Hello Ryan and Mr.Ghofrani")
}
func apisPrint(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"Apis":[]string{
			"GET   /hello",
			"POST   /print",
			"DELETE  /stop",
			"POST  /print",
		},
	})
}
func auth(c *gin.Context) {
	authHeader:=c.GetHeader("Authorization")
	if authHeader!="test"{
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Next()
}
func msg(c *gin.Context) {
	urlParam := c.PostForm("msg")
	if urlParam != "" {
		c.String(http.StatusOK, "Your Message is: %s", urlParam)
	}
		var jsonBody map[string]interface{}
		err:=c.ShouldBindJSON(&jsonBody);
	 if jsonBody!=nil{
		if  err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,jsonBody)
	}
	
 }
// func printJSON(c *gin.Context) {
// 	var jsonBody map[string]interface{}
// 	if err:=c.ShouldBindJSON(&jsonBody); err!=nil {
// 		c.JSON(http.StatusBadRequest,gin.H{
// 			"error":err.Error(),
// 		})
// 	}
// 	c.JSON(http.StatusOK,jsonBody)
// }
func main() {
router := gin.Default()
router.Use(auth)
router.GET("/hello", Hello)
router.GET("/", apisPrint)
router.POST("/print",msg )


router.DELETE("/stop", func(c *gin.Context) {
  os.Exit(0)
 })
router.Run(":8080")

}