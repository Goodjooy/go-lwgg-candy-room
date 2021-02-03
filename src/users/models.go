package users

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model

	Name         string `gorm:"type:varchar(32);not null" admin:"name:用户名;type:text"`
	EmailAddress string `gorm:"type:varchar(128)" admin:"name:邮箱;type:email"`

	UUID     string `gorm:"type:varchar(36);not null" admin:"autoG:uuidG;type:text"`
	PassWord string `gorm:"type:varchar(32);not null" admin:"name:用户密码哈希;256G;type:password"`

	AccountLevel uint8 `gorm:"not null;default:0" admin:"type:text"`
}

type UserInfo struct {
	gorm.Model

	PersonalSign string `gorm:"size:256"`
	PhoneNumber  string `grom:"size:32"`

	Sex byte

	Address string `gorm:"size:256"`

	User UserModel
	UserID uint
}

type UserImage struct {
	gorm.Model

	ImgPath string 	`gorm:"size:64"`

	User UserModel
	UserID uint
}