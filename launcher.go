package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go-lwgg-candy-room/src/admin"
	"go-lwgg-candy-room/src/goods"
	"go-lwgg-candy-room/src/index"
	"go-lwgg-candy-room/src/users"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//TemplatePath 模板文件夹
const TemplatePath = "./templates/**/*"

//StaticPath 静态文件文件夹
const StaticPath = "./static/**/*"

//MediaPath 媒体文件文件夹
const MediaPath = "./medias/**/**/*"

//SQLUser 数据库用户
const SQLUser = "root"

//SQLPasswd 数据库密码
const SQLPasswd = "wyq020222"

//DatabaseName 使用数据库名称
const DatabaseName = "marker_holder"

func main() {
	db, isOK := sqlConnection(SQLUser, SQLPasswd, DatabaseName)
	if !isOK {
		return
	}

	server := gin.Default()
	server.LoadHTMLGlob(TemplatePath)
	server.Static("/static", StaticPath)

	adminter := admin.NewAdminManager(db)
	adminter.AsignApplication(server, db)

	//db.Create(&admin.SuperUser{UUID: manage.UUIDGenerate(), Email: "964413011@qq.com", Passwd: manage.DateSHA256Hash("wyq020222")})

	adminter.PushApplication(index.NewIndexApplication(db).AsignApplication(server, db))
	adminter.PushApplication(users.NewUserApplication(db).AsignApplication(server, db))
	adminter.PushApplication(goods.NewGoodsApplication(db).AsignApplication(server, db))
	
	server.Run(":8081")

	defer db.Close()
}

func sqlConnection(userID, password, dbName string) (*gorm.DB, bool) {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charsetutf8&parseTime=True&loc=Local", userID, password, dbName))
	if err != nil {
		fmt.Printf("数据库连接失败，%s", err.Error())
		return nil, false
	}
	return db.Debug(), true
}
