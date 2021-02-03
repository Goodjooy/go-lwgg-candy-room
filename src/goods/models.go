package goods

import (
	"go-lwgg-candy-room/src/users"
	"github.com/jinzhu/gorm"
)


type GoodsInfo struct {
	gorm.Model
	//商品名称
	Name string `gorm:"size:64" admin:"name:商品名称"`
	ImagePath string `gorm:"size:256"`
	//商品价格
	Price float64 `admin:"name:商品价格"`

	RemainCount int

	Master users.UserModel 
	MasterID int
}
// GoodsBags 购物车，外键1->user,外键2—>goods
type GoodsBag struct {
	gorm.Model

	User users.UserModel
	UserID int

	Goods GoodsInfo
	GoodsID int

	Count int
}