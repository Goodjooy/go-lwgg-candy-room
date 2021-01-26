package admin

//SuperUser admin users
type SuperUser struct {
	UUID string `gorm:"size:36;primary_key;not null" admin:"name:ID"`

	Email string `gorm:"size:128;not null"  admin:"type:email"`
	Passwd string `gorm:"size:32;not null"`
}
