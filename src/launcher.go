package main

import (
	"github.com/gin-gonic/gin"

	"go-lwgg-candy-room/src/index"
)




func main() {
	server := gin.Default()

	index.NewIndexApplication().AsignApplication(server)

	server.Run()
}

