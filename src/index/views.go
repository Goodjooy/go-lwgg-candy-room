package index

import (
	"net/http"

	"go-lwgg-candy-room/src/manage"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func newMainPage(db *gorm.DB) manage.Viewer {
	pv := manage.NewViewer("/", db)
	pv.AsgnMethod(manage.GET)

	pv.AsignHandle(func(c *gin.Context) {
		var candys []candy
		candys = append(candys, candy{Img: "HTTPS://g-search1.alicdn.com/img/bao/uploaded/i4/imgextra/i1/24768522/O1CN01Y17Src2Cp7vWgF5Sv_!!0-saturn_solar.jpg_250x250.jpg_.webp",
			Name: "好耶", Model: gorm.Model{ID: 1}})
		candys = append(candys, candy{Img: `https://img.alicdn.com/imgextra/i1/11170294/O1CN01whOdq81E2h1BAqTry_!!0-saturn_solar.jpg_468x468q75.jpg`,
			Name: `英国女装新款性感透视优雅气质女人味翻领黑色长袖连衣裙大摆长裙`, Model: gorm.Model{ID: 2}})
		c.HTML(http.StatusOK, "index_page.html", gin.H{
			"candys": candys,
		})
	})
	return pv
}
