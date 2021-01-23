package admin

import (
	"fmt"
	"go-lwgg-candy-room/src/manage"
	"reflect"
	"strconv"
	"strings"

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

func newInfoMationViewer(db *gorm.DB, admin *AdminApplication) manage.Viewer {
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

					modelValue := reflect.ValueOf(targetModel).Elem()

					var values []feildData
					for _, v := range colNames[1:] {
						feild := modelValue.FieldByName(v)
						t := strings.ToLower(feild.Type().Name())
						var value string
						var InputType string = "text"
						if strings.HasPrefix(t, "uint") {
							value = strconv.Itoa(int(feild.Uint()))
						} else if strings.HasPrefix(t, "int") {
							value = strconv.Itoa(int(feild.Int()))
						} else if strings.HasPrefix(t, "time") {
							value = feild.String()
							InputType = "date"
						} else {
							value = feild.String()
						}

						values = append(values, feildData{
							Name:        feild.Type().Name(),
							InputType:   InputType,
							DefaultData: value,
							MaxLen:      36,
						})
					}
				}
			}
		}
	})

	return v

}

func getStringFormatcolsData(targetModel interface{}) {

}
