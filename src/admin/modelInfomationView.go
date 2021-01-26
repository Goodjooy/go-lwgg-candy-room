package admin

import (
	"fmt"
	"go-lwgg-candy-room/src/admin/modelloader"
	"go-lwgg-candy-room/src/manage"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type feildData struct {
	Name        string
	InputType   string
	MaxLen      uint
	DefaultData string
}

const ignoreFelid = "Model"

const (
	pAppName        = "appName"
	pModelName      = "modelName"
	pPrimaryKeyName = "pk"
)

func newInfomationViewer(db *gorm.DB, admin *AdminApplication) manage.Viewer {
	v := manage.NewViewer("/infomation/:appName/:modelName/:pk", db)
	v.AsgnMethod(manage.GET)

	v.AsignHandle(func(c *gin.Context) {
		appName := c.Param(pAppName)
		modelName := c.Param(pModelName)
		pkValue := c.Param(pPrimaryKeyName)

		if loginStatueCheck(c, db) {
			if c.Request.Method == manage.GET {
				isOK, targetModel := modelFinding(appName, modelName, c, admin)
				if isOK {
					_, colNames := getPrimanyKeyFirstColCols(targetModel)
					pkName := colNames[0]

					tableName := getTableName(targetModel)
					db.Raw(fmt.Sprintf("Select *from %s where %s=%s", tableName, pkName, pkValue)).Scan(targetModel)

					//modelValue := reflect.ValueOf(targetModel).Elem()

					Res:=modelloader.NewModel(targetModel,appName)

					c.HTML(http.StatusOK,"model_info.html",Res.HtMLTemplateData())
				}
			}
		}
	})

	return v

}

func getStringFormatcolsData(targetModel interface{}) {

}
