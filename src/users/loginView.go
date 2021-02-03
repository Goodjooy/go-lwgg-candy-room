package users

import (
	"fmt"
	"net/http"

	"go-lwgg-candy-room/src/manage"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const normalUser = 0
const bessinessUser = 1
const normalUserName = "普通用户"
const bessinessUserName = "商家"

const uidCookie = "__uuid"
const passwdHashCookie = "__ps_hash__"

const exitFlage = "-1"

const userMainPagePathFormat = "%s/users/%s/mainPage"

type userLogin struct {
	EmailAddress string `form:"email"`
	Passwd       string `form:"passwd"`
}

func newUserLoginView(db *gorm.DB) manage.Viewer {

	v := manage.NewViewer("/login", db)
	v.AsgnMethod(manage.GET, manage.POST)

	v.AsignHandle(func(c *gin.Context) {
		var users []UserModel

		if c.Request.Method == manage.POST {

			var user userLogin
			var userModel UserModel

			err := c.ShouldBind(&user)

			if err == nil {
				passwdHash := manage.DateSHA256Hash(user.Passwd)
				//search user
				db.Where(&UserModel{EmailAddress: user.EmailAddress, PassWord: passwdHash}).Find(&users)

				if len(users) != 0 {
					userModel = users[0]

					c.SetCookie(uidCookie, userModel.UUID, 3600, "/", "", false, true)
					c.SetCookie(passwdHashCookie, userModel.PassWord, 3600, "/", "", false, true)

					c.Redirect(http.StatusFound,
						fmt.Sprintf(userMainPagePathFormat,
							appURLRoot,
							userModel.Name))
				} else {
					c.Redirect(http.StatusFound, "/user/login?info=邮箱或者密码错误")
				}
			}
		} else if c.Request.Method == manage.GET {
			//check cookie
			isOK, user := CheckLogin(c, db, false)
			if isOK {
				c.Redirect(http.StatusFound, fmt.Sprintf(userMainPagePathFormat, appURLRoot, user.Name))
				return

			}

			c.HTML(http.StatusOK, "login_page.html", gin.H{
				"targetURL": appURLRoot + "/login",
				"info":      c.DefaultQuery("info", "")})
		} else {
			c.String(http.StatusNotFound, "不支持")
		}
	})

	return v
}

//CheckLogin 检查用户登录状态
func CheckLogin(c *gin.Context, db *gorm.DB, autoRedirect bool) (bool, UserModel) {
	var users []UserModel
	var userModel UserModel

	uuid, errUID := c.Cookie(uidCookie)
	password, errpaswd := c.Cookie(passwdHashCookie)
	if errUID == nil && errpaswd == nil && password != exitFlage {

		db.Where(&UserModel{UUID: uuid, PassWord: password}).Find(&users)

		var info UserInfo

		
		if len(users) >= 1 {
			userModel = users[0]
			info.User=userModel
			info.Sex=2
			db.Where(&info).FirstOrCreate(&info)

			return true, userModel
		}
	}
	if autoRedirect {
		c.Redirect(http.StatusFound, "/user/login?info=请先登录")
	}
	return false, UserModel{}
}
