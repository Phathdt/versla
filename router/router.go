package router

import (
	"github.com/gin-gonic/gin"
	"versla/app"
	"versla/modules/book/booktransport"
)

func New(app *app.App) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", ping)

		bookRoute := v1.Group("/books")
		{
			bookRoute.GET("/:id", booktransport.GetBook(app))
			bookRoute.POST("", booktransport.CreateBook(app))
		}
	}

	_ = r.Run()

	return r
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
