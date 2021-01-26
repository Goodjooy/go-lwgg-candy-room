package admin

import (
	"fmt"
	"go-lwgg-candy-room/src/manage"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const (
	modelNameH       = "modelName"
	modelColsH       = "modelCols"
	modelInfoH       = "modelInfo"
	appNameH         = "appName"
	primaryKeyIndexH = "pi"
)
const letterSplit = `[A-Z][a-z0-9]+`
const gormModleName = "Model"

type data struct {
	DataList []string
	InfoURL  string
}

func newModelManageView(db *gorm.DB, admin *AdminApplication) manage.Viewer {
	v := manage.NewViewer("/model/:appName/:modelName", db)
	v.AsgnMethod(manage.GET)

	v.AsignHandle(func(c *gin.Context) {
		if loginStatueCheck(c, db) {
			appName := c.Param("appName")
			modelName := strings.ToLower(c.Param("modelName"))

			if c.Request.Method == manage.GET {
				isOK, targetModel := modelFinding(appName, modelName, c, admin)
				if isOK {

					pkIndex, colNames := getPrimanyKeyFirstColCols(targetModel)
					colNames = append(colNames, "操作")

					var modelList []data

					MName := getTableName(targetModel)

					rows, err := db.Raw(fmt.Sprintf("select * from %s;", MName)).Rows()
					//var data []interface{};
					//db.Debug(). sca
					if err != nil {

					} else {
						defer rows.Close()
						for rows.Next() {
							db.ScanRows(rows, targetModel)
							md := targetModel

							v := reflect.ValueOf(md).Elem()
							var temp []string
							if colNames[0] == "id" && v.FieldByName(gormModleName).IsValid() {
								Si := v.Field(pkIndex)
								model := Si.Interface()
								mv := reflect.ValueOf(model).FieldByName("ID").Uint()
								temp = append(temp, strconv.Itoa(int(mv)))
							} else {
								temp = append(temp, v.FieldByName(colNames[0]).String())
							}

							for _, fieldName := range colNames[1:len(colNames)-1] {

								v.FieldByName(fieldName)
								value := v.FieldByName(fieldName).String()
								if strings.HasPrefix(value, "<uint") {
									value = strconv.Itoa(int(v.FieldByName(fieldName).Uint()))
								}
								temp = append(temp, value)
							}
							modelList = append(modelList, data{DataList: temp, InfoURL: fmt.Sprintf("/%s/%s/%s",
								appName, reflect.TypeOf(targetModel).Elem().Name(), temp[0])})
						}
						c.HTML(http.StatusOK, "model_detial.html", gin.H{
							modelNameH: modelName,
							appNameH:   appName,
							modelColsH: colNames,
							modelInfoH: modelList,
						})
					}

				}

			}
		}
	})

	return v

}

func modelFinding(appName, modelName string, c *gin.Context, admin *AdminApplication) (bool, interface{}) {
	var targetModel interface{} = nil
	models := admin.appsModel[appName]
	for _, v := range models {
		t := reflect.TypeOf(v)
		Name := strings.ToLower(t.Elem().Name())
		if Name == strings.ToLower(modelName) {
			targetModel = v
			break
		}
	}
	if targetModel == nil {
		c.String(http.StatusNotFound, "the model <%s> in app %s not found", modelName, appName)
		return false, nil
	}
	return true, targetModel

}

func getPrimanyKeyFirstColCols(targetModel interface{}) (int, []string) {
	r := reflect.TypeOf(targetModel).Elem()
	listCols := r.NumField()
	var colsName []string
	primaryKeyIndex := -1

	for i := 0; i < listCols; i++ {
		f := r.Field(i)
		var name string
		if f.Name == "Model" {
			name = "id"
			primaryKeyIndex = 0
		} else {
			name = f.Name
		}
		colsName = append(colsName, name)

		if primaryKeyIndex == -1 {
			ormTag := f.Tag.Get("gorm")
			isOK, _ := regexp.MatchString(`^.*?primary_key.*?$`, ormTag)
			if isOK {
				primaryKeyIndex = i
			}
		}
	}

	var colNames []string
	colNames = append(colNames, colsName[primaryKeyIndex])
	if primaryKeyIndex > 0 {
		colNames = append(colNames, colsName[:primaryKeyIndex]...)
	}
	if primaryKeyIndex < len(colsName)-1 {
		colNames = append(colNames, colsName[primaryKeyIndex+1:]...)
	}
	return primaryKeyIndex, colNames
}
func getTableName(targetModel interface{}) string {
	modelName := reflect.TypeOf(targetModel).Elem().Name()
	pattern := regexp.MustCompile(letterSplit)
	splits := pattern.FindAllString(modelName, -1)
	MName := strings.Join(splits, "_")
	MName = strings.ToLower(MName + "s")

	return MName
}
