package booktransport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"versla/app"
	"versla/modules/book/bookhdl"
	"versla/modules/book/bookmodel"
	"versla/modules/book/bookrepo"
	"versla/modules/book/bookstorage"
	"versla/sdk/sdkcm"
)

func UpdateBook(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := sdkcm.StrToUint(c.Param("id"))

		if err != nil {
			panic(err)
		}

		var data bookmodel.BookUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := app.DB

		storage := bookstorage.NewBookSQLStorage(db)
		repo := bookrepo.NewUpdateBookRepo(storage)
		hdl := bookhdl.NewUpdateBookHdl(repo)

		book, err := hdl.Response(c.Request.Context(), id, &data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse(book))
	}
}
