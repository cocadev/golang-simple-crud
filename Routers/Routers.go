package Routers

import (
	"net/http"
	"../Controllers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := r.Group("/api")
	{
		api.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
			  "message": "Thanks my go api",
			})
		})

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
	}

	return r
}
