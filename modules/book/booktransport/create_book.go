package booktransport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"versla/app"
	"versla/modules/book/bookhdl"
	"versla/modules/book/bookmodel"
	"versla/modules/book/bookrepo"
	"versla/modules/book/bookstorage"
)

func CreateBook(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data bookmodel.BookCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		DB := app.DB

		storage := bookstorage.NewBookSQLStorage(DB)
		repo := bookrepo.NewCreateBookRepo(storage)
		hdl := bookhdl.NewCreateBookHdl(repo)

		ID, err := hdl.Response(c.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
			"id":  ID,
		})
	}
}
