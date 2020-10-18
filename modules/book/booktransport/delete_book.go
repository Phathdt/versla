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

func DeleteBook(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := sdkcm.StrToUint(c.Param("id"))

		if err != nil {
			panic(err)
		}

		DB := app.DB

		storage := bookstorage.NewBookSQLStorage(DB)
		repo := bookrepo.NewDeleteBookRepo(storage)
		hdl := bookhdl.NewDeleteBookHdl(repo)

		err = hdl.Response(c.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse("ok"))
	}
}
