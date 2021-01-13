package manage

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

/*
Application 后端的每个单独应用
*/
type Application struct {
	URLPattern string

	viewers           []Viewer
	viewerURLPatterns []string

	models []interface{}
}

func NewApplication(URLPattern, templatesPath, staticPath string) Application {
	app := Application{URLPattern: URLPattern}

	return app
}

func methodLimitaion(f func(c *gin.Context), supportMethods []string) gin.HandlerFunc {
	temp := func(c *gin.Context) {
		isSupport := false
		method := c.Request.Method
		for _, v := range supportMethods {
			if v == method {
				isSupport = true
				break
			}
		}

		if isSupport {
			f(c)
		} else {
			c.String(http.StatusNotFound, fmt.Sprintf("the method \"%s\" not support", method))
		}
	}

	return temp
}

func (app*Application)AsignModels(models...interface{}){
	app.models=append(app.models, models...)
}

func (app *Application) AsignViewer(viewers ...Viewer) {
	for _, viewer := range viewers {
		viewerURLPattern := viewer.URLPattern
		exist := false

		for _, pattern := range app.viewerURLPatterns {
			if pattern == viewerURLPattern {
				exist = true
			}
		}

		if !exist {
			app.viewers = append(app.viewers, viewer)
			app.viewerURLPatterns = append(app.viewerURLPatterns, viewerURLPattern)
		}
	}
}

func (app Application) AsignApplication(server *gin.Engine, db *gorm.DB) {
	group := server.Group(app.URLPattern)
	viewers := app.viewers
	for _, v := range viewers {
		supportMethod := (v).SupportMethods
		group.Any((v).URLPattern, methodLimitaion((v).handle, supportMethod))
	}

	for _, model := range app.models {
		if !db.HasTable(model) {
			db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(model)
		}
	}
}
