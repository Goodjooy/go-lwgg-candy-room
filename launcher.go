package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go-lwgg-candy-room/src/admin"
	"go-lwgg-candy-room/src/goods"
	"go-lwgg-candy-room/src/index"
	"go-lwgg-candy-room/src/manage"
	"go-lwgg-candy-room/src/users"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, isOK := sqlConnection(manage.SQLUser, manage.SQLPasswd, manage.DatabaseName)
	if !isOK {
		return
	}

	server := gin.Default()
	server.LoadHTMLGlob(manage.TemplatePath)
	server.Static(manage.StaticURLRoot, manage.StaticPath)
	server.Static(manage.MediaURLRoot, manage.MediaPath)

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
