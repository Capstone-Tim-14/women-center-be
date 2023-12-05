package schema

type BankMethod struct {
	Id       uint   `gorm:"primaryKey"`
	Title    string `gorm:"type:varchar(50)"`
	BankCode string `gorm:"type:enum('bni','mandiri','bri','bca')"`
	Image    string `gorm:"varchar(255)"`
}

type WalletMethod struct {
	Id         uint   `gorm:"primaryKey"`
	Title      string `gorm:"type:varchar(50)"`
	WalletCode string `gorm:"type:enum('gopay','qris')"`
	Image      string `gorm:"varchar(255)"`
}
