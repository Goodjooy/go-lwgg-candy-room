package manage

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	OPTIONS = "OPTIONS"
)

type Viewer struct {
	URLPattern     string
	SupportMethods []string

	handle []gin.HandlerFunc
}

func NewViewer(URLPattern string,db *gorm.DB) Viewer {
	return Viewer{URLPattern: URLPattern}
}

func (v *Viewer) AsgnMethod(methods ...string) {
	for _, method := range methods {
		v.SupportMethods = append(v.SupportMethods, method)
	}
}

func (v *Viewer) AsignHandle(handle... gin.HandlerFunc) {
	v.handle = append(v.handle, handle...)
}
