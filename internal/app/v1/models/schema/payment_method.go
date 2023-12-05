package schema

type PaymentMethod struct {
	Id    uint   `gorm:"primaryKey"`
	Title string `gorm:"type:varchar(50)"`
}
