package admin

import (
	"go-lwgg-candy-room/src/manage"

	"github.com/jinzhu/gorm"
)

//管理员部分
const appRootURL="/admin"

type Application interface {
	GetAllModels()[]interface{}
}

type AdminApplication struct {
	manage.Application

	appsModel map[string][]interface{}
}

func NewAdminManager(db *gorm.DB) AdminApplication {
	app:=AdminApplication{}
	app.Application=manage.NewApplication("/admin","","")

	app.AsignModels(&SuperUser{})


	app.appsModel=make(map[string][]interface{})
	return app
}

