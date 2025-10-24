package entity

type Admin struct {
	Id       uint `gorm:"primaryKey;autoIncrement"`
	Nama     string
	Email    string
	Password string
	CreateAt string
	UpdateAt string
}
