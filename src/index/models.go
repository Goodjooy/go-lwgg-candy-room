package index

import "github.com/jinzhu/gorm"

type candy struct {
	gorm.Model

	Img  string
	Name string
}