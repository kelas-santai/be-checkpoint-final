package entity

type Merchant struct {
	Id       uint `gorm:"primaryKey;autoIncrement"`
	Nama     string
	Email    string
	Password string
	Lokasi   string
	Foto     string
	NoTelpon string
	CreateAt string
	UpdateAt string
}
