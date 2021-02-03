package index

import (
	"net/http"

	"go-lwgg-candy-room/src/goods"
	"go-lwgg-candy-room/src/manage"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func newMainPage(db *gorm.DB) manage.Viewer {
	pv := manage.NewViewer("/", db)
	pv.AsgnMethod(manage.GET)

	pv.AsignHandle(func(c *gin.Context) {
		var goods []goods.GoodsInfo

		db.Find(&goods)
		c.HTML(http.StatusOK, "index_page.html", goods)
	})
	return pv
}
