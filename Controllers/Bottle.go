package Controllers

import (
	"../ApiHelpers"
	"../Models"
	
	"github.com/gin-gonic/gin"
)

func ListBottle(c *gin.Context) {
	var bottle []Models.Bottle
	err := Models.GetAllBottle(&bottle)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, bottle)
	} else {
		ApiHelpers.RespondJSON(c, 200, bottle)
	}
}

func GetOneBottle(c *gin.Context) {
	id := c.Params.ByName("id")
	var bottle Models.Bottle
	err := Models.GetOneBottle(&bottle, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, bottle)
	} else {
		ApiHelpers.RespondJSON(c, 200, bottle)
	}
}