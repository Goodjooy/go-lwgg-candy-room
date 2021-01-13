package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go-lwgg-candy-room/src/index"
	"go-lwgg-candy-room/src/users"
)
import _"github.com/jinzhu/gorm/dialects/mysql"

//TemplatePath 模板文件夹
const TemplatePath = "../templates/**/*"

//StaticPath 静态文件文件夹
const StaticPath = "../static/**/*"

//MediaPath 媒体文件文件夹
const MediaPath = "../medias/**/**/*"

//SQLUser 数据库用户
const SQLUser = "root"

//SQLPasswd 数据库密码
const SQLPasswd = "wyq020222"

//DatabaseName 使用数据库名称
const DatabaseName = "marker_holder"

func main() {
	db, isOK := SqlConnection(SQLUser, SQLPasswd, DatabaseName)
	if !isOK {
		return
	}

	server := gin.Default()
	server.LoadHTMLGlob(TemplatePath)
	server.Static("/static", StaticPath)

	index.NewIndexApplication(db).AsignApplication(server,db)
	users.NewUserApplication(db).AsignApplication(server,db)

	server.Run()

	defer db.Close()
}

func SqlConnection(userID, password, dbName string) (*gorm.DB, bool) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charsetutf8&parseTime=True&loc=Local", userID, password, dbName))
	if err != nil {
		fmt.Printf("数据库连接失败，%s", err.Error())
		return nil, false
	}
	return db, true
}
