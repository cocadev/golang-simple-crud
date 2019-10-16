package Controllers

import (
	"../Models"
	"../ApiHelpers"

	"github.com/gin-gonic/gin"
) 

func ListJoke(c *gin.Context) {
	var joke []Models.Joke
	err := Models.GetAllJoke(&joke)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, joke)
	} else {
		ApiHelpers.RespondJSON(c, 200, joke)
	}
}

func AddNewJoke(c *gin.Context) {
	var joke Models.Joke
	c.BindJSON(&joke)
	err := Models.AddNewJoke(&joke)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, joke)
	} else {
		ApiHelpers.RespondJSON(c, 200, joke)
	}
}

func GetOneJoke(c *gin.Context) {
	id := c.Params.ByName("id")
	var joke Models.Joke
	err := Models.GetOneJoke(&joke, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, joke)
	} else {
		ApiHelpers.RespondJSON(c, 200, joke)
	}
}

func PutOneJoke(c *gin.Context) {
	var joke Models.Joke
	id := c.Params.ByName("id")
	err := Models.GetOneJoke(&joke, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, joke)
	}
	c.BindJSON(&joke)
	err = Models.PutOneJoke(&joke, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, joke)
	} else {
		ApiHelpers.RespondJSON(c, 200, joke)
	}
}

func DeleteJoke(c *gin.Context) {
	var joke Models.Joke
	id := c.Params.ByName("id")
	err := Models.DeleteJoke(&joke, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, joke)
	} else {
		ApiHelpers.RespondJSON(c, 200, joke)
	}
}
