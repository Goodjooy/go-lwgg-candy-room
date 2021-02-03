package goods

import (
	"go-lwgg-candy-room/src/manage"
	"go-lwgg-candy-room/src/users"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type myBagHTMLData struct {
	User users.UserModel
	Goods []GoodsBag
}

func newUserBagViewer(db *gorm.DB)manage.Viewer{
	v:=manage.NewViewer("/my-bag",db)

	v.AsgnMethod(manage.GET)

	v.AsignHandle(func(c *gin.Context) {
		if c.Request.Method==manage.GET{
			isOK,user:=users.CheckLogin(c,db,true)
			if isOK{
				HTMLData:=myBagHTMLData{User: user}
				exampleBag:=GoodsBag{
					UserID:int(user.Model.ID),
				}
				var bags []GoodsBag

				db.Where(&exampleBag).Find(&bags)

				for _, bag := range bags {
					exampleGoodInfo:=GoodsInfo{Model: gorm.Model{ID: uint(bag.GoodsID)}}
					var goods []GoodsInfo
					db.Where(&exampleGoodInfo).Find(&goods)
					if len(goods)>=1{
						bag.Goods=goods[0];
						HTMLData.Goods=append(HTMLData.Goods, bag)
					}

				}
				c.HTML(http.StatusOK,"my_bag_page.html",HTMLData)
				return
			}
		}

	})

	return v
}