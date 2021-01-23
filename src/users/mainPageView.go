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
		var userInfo UserModel
		var users []UserModel
		if c.Request.Method == manage.GET {
			t := gin.H{}
			userName := c.Param("userName")
			db.Where(&UserModel{Name: userName}).Find(&users)

			uuid, erruuid := c.Cookie(uidCookie)
			pass, errpass := c.Cookie(passwdHashCookie)

			if len(users) >= 1 {

				userInfo = users[0]

				var userLevel string
				switch userInfo.AccountLevel {
				case normalUser:
					userLevel = normalUserName
				case bessinessUser:
					userLevel = bessinessUserName
				default:
					userLevel = "unknow"
				}
				if errpass == nil && erruuid == nil && pass != exitFlage {
					if userInfo.PassWord == pass && userInfo.UUID == uuid {

						t = gin.H{
							userNameKey:        userInfo.Name,
							userLevelKey:       userLevel,
							userShoppingBagKey: "/",
							logoutKey:          "/user/logout",
						}
					}
				} else {
					t = gin.H{
						userNameKey:        userInfo.Name,
						userLevelKey:       userLevel,
						userShoppingBagKey: "/",
						logoutKey:          "/",
					}
				}
				c.HTML(http.StatusOK, "user_main_page.html", t)
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
