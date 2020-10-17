package booktransport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"versla/app"
	"versla/modules/book/bookhdl"
	"versla/modules/book/bookrepo"
	"versla/modules/book/bookstorage"
)

func GetBook(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		unitId, err := strconv.ParseUint(id, 10, 16)

		if err != nil {
			panic(err)
		}

		DB := app.DB

		storage := bookstorage.NewBookSQLStorage(DB)
		repo := bookrepo.NewGetBookRepo(storage)
		hdl := bookhdl.NewGetBookHdl(repo)

		book, err := hdl.Response(c.Request.Context(), unitId)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": book,
		})
	}
}
