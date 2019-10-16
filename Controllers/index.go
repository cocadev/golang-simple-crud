
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
