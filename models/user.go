package models

type User struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"type:varchar(255)" json:"name"`
	Uname string `gorm:"type:varchar(100);unique" json:"uname"`
	Pass  string `gorm:"type:varchar(100)" json:"pass"`
}
