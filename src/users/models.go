package users

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model

	Name         string `gorm:"type:varchar(32);not null"`
	EmailAddress string `gorm:"type:varchar(128)"`

	UUID     string `gorm:"type:varchar(36);not null"`
	PassWord string `gorm:"type:varchar(32);not null"`

	AccountLevel uint8 `gorm:"not null;default:0"`
}
