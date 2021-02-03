package goods

import (
	"go-lwgg-candy-room/src/manage"
	"github.com/jinzhu/gorm"
)

const appRootURL="/good"
func NewGoodsApplication(db *gorm.DB) manage.Application{
	app:=manage.NewApplication("/goods","goods","")

	app.AsignModels(&GoodsInfo{})
	app.AsignModels(&GoodsBag{})

	app.AsignViewer(newUserBagViewer(db))
	app.AsignViewer(newGoodsInfoViewer(db))

	return app;
}