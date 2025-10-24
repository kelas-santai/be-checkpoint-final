package entity

type User struct {
	Id          uint `gorm:"primaryKey;autoIncrement"`
	Nama        string
	Email       string
	Password    string
	KodeReferal string
	NoTelpon    string
	CreateAt    string
	UpdateAt    string
}
