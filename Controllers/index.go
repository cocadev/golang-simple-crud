
package Controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"message": "Thanks my go api",
	 })
}

func Welcome(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H {
	  "message":"Welcome",
	})
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}