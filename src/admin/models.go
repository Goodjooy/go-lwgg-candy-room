package admin


type SuperUser struct {
	UUID string `gorm:"size:36;primary_key;not null"`

	email string `gorm:"size:128;not null"`
	passwd string `gorm:"size:32;not null"`
}
