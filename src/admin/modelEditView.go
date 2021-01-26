package admin

import (
	"go-lwgg-candy-room/src/manage"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func newModelEditViewer(db *gorm.DB,admin* AdminApplication)manage.Viewer{
	v:=manage.NewViewer("/edit/:appName/:modelName",db)
	v.AsgnMethod(manage.POST)

	v.AsignHandle(func(c *gin.Context) {
		appName:=c.Param("appName")
		modelName:=c.Param("modelName")
		//pk:=c.DefaultQuery("pk","")


		if loginStatueCheck(c,db){
			modelFinding(appName,modelName,c,admin)
		}
	})

	return v
}