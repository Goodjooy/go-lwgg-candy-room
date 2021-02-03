package users

import (
	"go-lwgg-candy-room/src/manage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

const _PCOldPassword = "opw"
const _PCNewPassword = "np1"
const _PCNewPasswordCheck = "np2"

type passwordChange struct {
	Old      string `form:"opw"`
	New      string `form:"np1"`
	NewCheck string `form:"np2"`
}

func newPasswordChangeViewer(db *gorm.DB) manage.Viewer {
	v := manage.NewViewer("/password-change", db)

	v.AsgnMethod(manage.GET, manage.POST)

	v.AsignHandle(func(c *gin.Context) {
		//check log in
		isOK, user := CheckLogin(c, db, true)
		if isOK {
			if c.Request.Method == manage.POST {
				var password passwordChange
				if c.MustBindWith(&password, binding.Form) == nil {
					if manage.DateSHA256Hash(password.Old) == user.PassWord &&
						password.New == password.NewCheck {
							user.PassWord=manage.DateSHA256Hash(password.New)
							db.Save(&user)
							c.Redirect(http.StatusFound,"/user/login?info=密码修改完成，请重新登录")
					}
				}
			}else if c.Request.Method==manage.GET {
				c.HTML(http.StatusOK,"password_change.html",nil)
			}
		}

	})

	return v
}
