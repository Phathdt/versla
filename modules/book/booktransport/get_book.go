package booktransport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"versla/app"
	"versla/modules/book/bookhdl"
	"versla/modules/book/bookrepo"
	"versla/modules/book/bookstorage"
	"versla/sdk/sdkcm"
)

func GetBook(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := sdkcm.StrToUint(c.Param("id"))

		if err != nil {
			panic(err)
		}

		DB := app.DB

		storage := bookstorage.NewBookSQLStorage(DB)
		repo := bookrepo.NewGetBookRepo(storage)
		hdl := bookhdl.NewGetBookHdl(repo)

		book, err := hdl.Response(c.Request.Context(), uint32(id))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse(book))
	}
}
