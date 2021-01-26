package admin

import (
	"go-lwgg-candy-room/src/manage"

	"github.com/jinzhu/gorm"
)

//管理员部分
const appRootURL = "/admin"
//Application app interface
type Application interface {
	GetAllModels() []interface{}
	GetAppName() string
}
//AdminApplication admin app
type AdminApplication struct {
	manage.Application

	appsModel map[string][]interface{}
}
//NewAdminManager creat admon app
func NewAdminManager(db *gorm.DB) AdminApplication {
	app := AdminApplication{}
	app.Application = manage.NewApplication("/admin", "admin", "")

	app.AsignModels(&SuperUser{})

	app.AsignViewer(newModelManageView(db, &app))
	app.AsignViewer(newMainPageView(db,&app))
	app.AsignViewer(newLoginPageView(db))
	app.AsignViewer(newInfomationViewer(db,&app))

	app.appsModel = make(map[string][]interface{})

	app.PushApplication(app)
	return app
}

//PushApplication add app under admin ctrl
func (admin *AdminApplication) PushApplication(apps ...Application) {
	for _, app := range apps {
		admin.appsModel[app.GetAppName()] = app.GetAllModels()
	}
}
