package entity

//User Ini Blom Di Buat CRUD nya
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

//User Product Ini Blom di buat juga
type UsersProduct struct {
	Id          uint `gorm:"primaryKey;autoIncrement"`
	IdUser      uint
	IdProduct   uint
	NoTraksaksi string
	CreateAt    string
	UpdateAt    string
}
