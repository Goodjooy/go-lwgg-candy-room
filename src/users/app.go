package users

import (
	"go-lwgg-candy-room/src/manage"
	"github.com/jinzhu/gorm"
)

const appURLRoot="/user"


func NewUserApplication(db *gorm.DB)manage.Application{
	app:=manage.NewApplication(appURLRoot,"user","")

	app.AsignModels(&UserModel{})
	app.AsignModels(&UserInfo{})
	app.AsignModels(&UserImage{})

	app.AsignViewer(newUserLoginView(db))
	app.AsignViewer(newSignupViewer(db))
	app.AsignViewer(newUserMainView(db))
	app.AsignViewer(newUserExitView(db))
	app.AsignViewer(newUserLevelUp(db))
	app.AsignViewer(newPasswordChangeViewer(db))
	app.AsignViewer(newPersonInfoEditViewer(db))
	app.AsignViewer(newUploadImageViewer(db))

	return app
}