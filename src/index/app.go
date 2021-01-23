package index

import (
	"github.com/jinzhu/gorm"
	"go-lwgg-candy-room/src/manage"
)

func NewIndexApplication(db *gorm.DB) manage.Application {
	app := manage.NewApplication("/", "index", "")

	app.AsignViewer(newMainPage(db))

	return app
}
