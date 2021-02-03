package users

import (
	"fmt"
	"go-lwgg-candy-room/src/manage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func newUploadImageViewer(db *gorm.DB) manage.Viewer {
	v := manage.NewViewer("/uploadImg", db)
	v.AsgnMethod(manage.GET, manage.POST)

	v.AsignHandle(func(c *gin.Context) {
		isOK, user := CheckLogin(c, db, true)
		if isOK {
			if c.Request.Method == manage.GET {
				c.HTML(http.StatusOK, "file_upload.html", gin.H{
					"targetURL":  appURLRoot + "/uploadImg",
					"acceptFile": ".png,.jpg",
				})
			} else if c.Request.Method == manage.POST {
				fileName := manage.UploadFileViewHandle(c, "file", "user")
				var img UserImage
				img.User = user
				db.Where(&img).FirstOrCreate(&img)

				img.ImgPath = fileName
				db.Save(&img)

				c.Redirect(http.StatusFound, fmt.Sprintf(userMainPagePathFormat, appURLRoot, user.Name))
			}
		}

	})

	return v
}
