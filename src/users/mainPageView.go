package users

import (
	"go-lwgg-candy-room/src/manage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	userNameKey        = "name"
	userLevelKey       = "level"
	userShoppingBagKey = "user_bag"
	logoutKey          = "user_logout"
)

func newUserMainView(db *gorm.DB) manage.Viewer {
	v := manage.NewViewer("/users/:userName/mainPage", db)
	v.AsgnMethod(manage.GET)

	v.AsignHandle(func(c *gin.Context) {
		if c.Request.Method == manage.GET {
			isOK, user := CheckLogin(c, db,true)
			if isOK {
				var info UserInfo
				info.User = user
				var img UserImage
				img.User = user

				db.Where(&info).FirstOrCreate(&info)
				db.Where(&img).FirstOrCreate(&img)

				c.HTML(http.StatusOK, "user_main_page.html", gin.H{
					"user": user,
					"info": info,
					"img":  img,
				})
			}

		}

	})
	return v
}
func newUserLevelUp(db *gorm.DB) manage.Viewer {
	v := manage.NewViewer("/upgrade", db)
	v.AsgnMethod(manage.POST)

	v.AsignHandle(func(c *gin.Context) {
		if c.Request.Method == manage.POST {

			isOK, _ := CheckLogin(c, db, true)
			if isOK {
				c.String(http.StatusForbidden, "abab?")
				//user.AccountLevel ^= 1
				//db.Save(&user)
				//c.Redirect(http.StatusForbidden,"/user/login")
			}
		}
	})

	return v
}

func newUserExitView(db *gorm.DB) manage.Viewer {
	v := manage.NewViewer("/logout", db)
	v.AsgnMethod(manage.POST)

	v.AsignHandle(func(c *gin.Context) {

		if c.Request.Method == manage.POST {
			pass, errpass := c.Cookie(passwdHashCookie)
			if errpass == nil && pass != exitFlage {
				c.SetCookie(passwdHashCookie, exitFlage, 3600, "/", "localhost", false, true)
				c.Redirect(http.StatusMovedPermanently, "/user/login")
			}

		}
	})
	return v
}
