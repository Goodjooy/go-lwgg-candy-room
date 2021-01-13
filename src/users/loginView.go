package users

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"go-lwgg-candy-room/src/manage"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const uidCookie = "__uuid"
const passwdHashCookie = "__ps_hash__"

const userMainPagePathFormat = "%s/%s/mainPage"

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
				passwdHash := passwdHash(user.Passwd)
				//search user
				db.Where(&UserModel{EmailAddress: user.EmailAddress, PassWord: passwdHash}).Find(&users)

				if len(users) != 0 {
					userModel = users[0]

					c.SetCookie(uidCookie, userModel.UUID, 3600, "/", "", false, true)
					c.SetCookie(passwdHashCookie, userModel.PassWord, 3600, "/", "", false, true)

					c.Redirect(http.StatusMovedPermanently,
						fmt.Sprintf(userMainPagePathFormat,
							appURLRoot,
							userModel.Name))
				} else {
					c.Redirect(http.StatusMovedPermanently, "/user/login")
				}
			}
		} else if c.Request.Method == manage.GET {
			//check cookie
			uuid, errUID := c.Cookie(uidCookie)
			password, errpaswd := c.Cookie(passwdHashCookie)
			if errUID == nil && errpaswd == nil {
				var userModel UserModel

				db.Where(&UserModel{UUID: uuid, PassWord: password}).Find(&users)

				if len(users) != 0 {
					userModel = users[0]

					c.Redirect(http.StatusMovedPermanently, fmt.Sprintf(userMainPagePathFormat, appURLRoot, userModel.Name))
					return
				}
			}

			c.HTML(http.StatusOK, "login_page.html", gin.H{})
		} else {
			c.String(http.StatusNotFound, "不支持")
		}
	})

	return v
}

const numStart='0'
const lowAlphaStart='a'
const upAlphaStart='A'
func passwdHash(passwd string) string {
	hash := sha256.Sum256([]byte(passwd))

	var hashRe []byte

	for _, v := range hash {
		t:=v%(10+26+26)
		if t<10 {
			t=t+numStart
		}else if t<26+10 {
			t=t+lowAlphaStart-10
		} else {
			t=t+upAlphaStart-10-26
		}
		hashRe = append(hashRe, t)
	}

	return string(hashRe)

}
