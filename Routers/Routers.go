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

		jokes := api.Group("/jokes")
		{
			jokes.GET("/", Controllers.ListJoke)
			jokes.GET("/like/:jokeID", Controllers.LikeOneJoke)
		}
	}

	return r
}
