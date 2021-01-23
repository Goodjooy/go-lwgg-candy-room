package admin

import (
	"fmt"
	"go-lwgg-candy-room/src/manage"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const adminUUIDCookieName = "__ad_uuid__cck__"
const adminPasswordHashCookieName = "__ad_pd__hash__"

const passwordExitSign = "-1"

type modelData struct {
	AppName    string
	ModelsName []Name
}
type Name struct {
	Name string
	URL  string
}

func NewMainPageView(db *gorm.DB, admin *AdminApplication) manage.Viewer {
	v := manage.NewViewer("/", db)

	v.AsgnMethod(manage.GET)

	v.AsignHandle(func(c *gin.Context) {
		if loginStatueCheck(c, db) {
			if c.Request.Method == manage.GET {
				var dataGroup []modelData
				for name, values := range admin.appsModel {
					data := modelData{AppName: name}
					for _, value := range values {
						modelName := reflect.TypeOf(value).Elem().Name()
						data.ModelsName = append(data.ModelsName, Name{Name: modelName, URL: fmt.Sprintf("/%s/%s", name, modelName)})
					}
					dataGroup = append(dataGroup, data)
				}
				c.HTML(http.StatusOK, "main_page.html", gin.H{
					"models": dataGroup,
				})
			}
		}

	})

	return v
}
func newLoginPageView(db *gorm.DB) manage.Viewer {
	v := manage.NewViewer("/login", db)
	v.AsgnMethod(manage.GET, manage.POST)

	v.AsignHandle(func(c *gin.Context) {
		if c.Request.Method == manage.POST {
			userID := c.PostForm("email-address")
			passwd := c.PostForm("passwd")
			pawd := manage.DateSHA256Hash(passwd)

			if userID != "" && passwd != "" {
				var users []SuperUser
				db.Where(&SuperUser{Email: userID, Passwd: pawd}).Find(&users)
				if len(users) > 0 {
					user := users[0]
					c.SetCookie(adminUUIDCookieName, user.UUID, 3600, "/", "", false, true)
					c.SetCookie(adminPasswordHashCookieName, user.Passwd, 3600, "/", "", false, true)
				}
			}
			c.Redirect(http.StatusMovedPermanently, "/admin/")
		} else {
			c.HTML(http.StatusOK, "login_page.html", gin.H{
				"targetURL": "/admin/login",
			})
		}
	})

	return v
}

func loginStatueCheck(c *gin.Context, db *gorm.DB) bool {
	//test code
	//return true;

	var superUsers []SuperUser

	//check cookie
	uuid, erruid := c.Cookie(adminUUIDCookieName)
	passHash, errpasswd := c.Cookie(adminPasswordHashCookieName)

	if errpasswd == nil && erruid == nil && passHash != passwordExitSign {
		db.Where(&SuperUser{UUID: uuid, Passwd: passHash}).Find(&superUsers)
		if len(superUsers) >= 1 {
			//render page
			return true
		}

	}
	c.Redirect(http.StatusMovedPermanently, appRootURL+"/login")
	return false
}
