package users

import (
	"go-lwgg-candy-room/src/manage"
	"github.com/jinzhu/gorm"
)

const appURLRoot="/user"


func NewUserApplication(db *gorm.DB)manage.Application{
	app:=manage.NewApplication(appURLRoot,"","")

	app.AsignModels(&UserModel{})

	app.AsignViewer(newUserLoginView(db))
	app.AsignViewer(newSignupViewer(db))
	app.AsignViewer(newUserMainView(db))
	app.AsignViewer(newUserExitView(db))
	
	return app
}