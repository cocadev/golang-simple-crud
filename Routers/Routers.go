package Routers

import (
	"../Controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := r.Group("/api")
	{
		api.GET("", Controllers.GetHome)
		api.GET("welcome", Controllers.Welcome)
		api.GET("hello", Controllers.Hello)

		bottles := api.Group("/bottle")
		{
			bottles.GET("/:id", Controllers.GetOneBottle)
			bottles.GET("", Controllers.ListBottle)
		}

		books := api.Group("/book")
		{
			books.GET("", Controllers.ListBook)
			books.POST("", Controllers.AddNewBook)
			books.GET("/:id", Controllers.GetOneBook)
			books.PUT("/:id", Controllers.PutOneBook)
			books.DELETE("/:id", Controllers.DeleteBook)
		}

		jokes := api.Group("/joke")
		{
			jokes.GET("", Controllers.ListJoke)
			jokes.POST("", Controllers.AddNewJoke)
			jokes.GET("/:id", Controllers.GetOneJoke)
			jokes.PUT("/:id", Controllers.PutOneJoke)
			jokes.DELETE("/:id", Controllers.DeleteJoke)
		}
	}

	return r
}
