package admin

import (
	"go-lwgg-candy-room/src/manage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const adminUUIDCookieName="__ad_uuid__cck__"
const adminPasswordHashCookieName="__ad_pd__hash__"

const passwordExitSign="-1"

func NewMainPageView(db *gorm.DB) manage.Viewer {
	v:=manage.NewViewer("/",db)

	v.AsgnMethod(manage.GET)

	v.AsignHandle(func(c *gin.Context) {
		if c.Request.Method==manage.GET{
			var superUsers []SuperUser;

			//check cookie
			uuid,erruid:=c.Cookie(adminUUIDCookieName)
			passHash,errpasswd:=c.Cookie(adminPasswordHashCookieName)

			if errpasswd ==nil &&erruid==nil && passHash!=passwordExitSign{
				db.Where(&SuperUser{UUID: uuid,passwd: passHash}).Find(&superUsers)
				if len(superUsers)>=1 {
					//render page
					return
				}

			}
			c.Redirect(http.StatusMovedPermanently,appRootURL+"/login")
		}

	})

	return v
}