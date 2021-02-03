package users

import (
	"fmt"
	"go-lwgg-candy-room/src/manage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

type userNewInfo struct {
	PersonalSign string `form:"sign"`
	PhoneNumber  string `form:"phone"`
	Sex          byte   `form:"sex"`
	Address      string `form:"address"`
}

func newPersonInfoEditViewer(db *gorm.DB) manage.Viewer {
	v := manage.NewViewer("/edit-info", db)
	v.AsgnMethod(manage.GET, manage.POST)

	v.AsignHandle(func(c *gin.Context) {
		isOk, user := CheckLogin(c, db, true)
		if isOk {
			if c.Request.Method == manage.POST {
				var userNewInfo userNewInfo

				if c.MustBindWith(&userNewInfo,binding.Form) == nil {
					var userInfo UserInfo
					userInfo.User=user
					db.Where(&userInfo).FirstOrCreate(&userInfo)

					userInfo.PersonalSign=userNewInfo.PersonalSign
					userInfo.PhoneNumber=userNewInfo.PhoneNumber
					userInfo.Sex=userNewInfo.Sex
					userInfo.Address=userNewInfo.Address

					db.Save(&userInfo)
					c.Redirect(http.StatusFound,fmt.Sprintf(userMainPagePathFormat,appURLRoot,user.Name))
				}
			}else if c.Request.Method==manage.GET {
				var userInfo UserInfo
				userInfo.User=user
				db.Where(&userInfo).FirstOrCreate(&userInfo)

				var sexGroup =[]string{"false","false","false"}
				sexGroup[int(userInfo.Sex)]="true"


				c.HTML(http.StatusOK,"info_change.html",gin.H{
					"user":userInfo,
					"sex":sexGroup,
				})
			}

		}
	})

	return v
}
