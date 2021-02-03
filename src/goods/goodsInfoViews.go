package goods

import (
	"go-lwgg-candy-room/src/manage"
	"go-lwgg-candy-room/src/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func newGoodsInfoViewer(db *gorm.DB) manage.Viewer {
	v := manage.NewViewer("/info/:goodsPK", db)
	v.AsgnMethod(manage.GET)

	v.AsignHandle(func(c *gin.Context) {
		goodsPK:=c.Param("goodsPK")
		pk,_:=strconv.Atoi(goodsPK)

		var goods GoodsInfo
		goods.ID=uint(pk)
		db.Where(&goods).First(&goods)

		var user users.UserModel
		user.ID=goods.Master.ID
		db.Where(&user).First(&user)

		goods.Master=user;

		c.HTML(http.StatusOK,"goods_info.html",goods)
	})

	return v
}
