package users

import (
	"fmt"
	"go-lwgg-candy-room/src/manage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type signupData struct {
	Name  string `form:"user-name"`
	Email string `form:"email"`

	passWd   string `form:"passwd"`
	passWdCk string `form:"passwd-ck"`

	AcLv string `form:"account-level"`
}

func newSignupViewer(db *gorm.DB) manage.Viewer {
	v := manage.NewViewer("/signup", db)
	v.AsgnMethod(manage.GET, manage.POST)

	v.AsignHandle(func(c *gin.Context) {
		if c.Request.Method == manage.POST {
			name := c.PostForm("user-name")
			passwd := c.PostForm("passwd")
			acLv := c.PostForm("account-level")

			var data signupData
			var dataGroup []UserModel
			var user UserModel

			err := c.ShouldBind(&data)

			if err == nil {
				db.Where(&UserModel{Name: data.Name}).Or(&UserModel{EmailAddress: data.Email}).Find(&dataGroup)
				if data.passWd == data.passWdCk && len(dataGroup) == 0 {
					//pass check
					accountLV, _ := strconv.Atoi(acLv)
					user = UserModel{Name: name,
						EmailAddress: data.Email,
						PassWord:     manage.DateSHA256Hash(passwd),
						UUID:         manage.UUIDGenerate(),
						AccountLevel: uint8(accountLV)}

					db.Create(&user)

					c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s/login",
						appURLRoot))
					return
				}
			}
			c.SetCookie(uidCookie, user.UUID, 3600, "/", "", false, true)
			c.SetCookie(passwdHashCookie, user.PassWord, 3600, "/", "", false, true)

			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("%s/signup", appURLRoot))
		} else if c.Request.Method == manage.GET {
			c.HTML(http.StatusOK, "signup_page.html", nil)
		}
	})

	return v
}

