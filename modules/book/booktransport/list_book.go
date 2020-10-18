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

func ListBook(app *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var p bookmodel.ListParams

		if err := c.ShouldBind(&p); err != nil && err.Error() != "EOF" {
			panic(err)
		}

		p.FullFill()

		DB := app.DB

		storage := bookstorage.NewBookSQLStorage(DB)
		repo := bookrepo.NewListBookRepo(storage)
		hdl := bookhdl.NewListBookHdl(repo)

		data, err := hdl.Response(c.Request.Context(), p.ListFilter, &p.Paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, sdkcm.ResponseWithPaging(data, p.ListFilter, p.Paging))
		//c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse(data))
	}
}
