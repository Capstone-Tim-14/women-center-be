package domain

import (
	"time"
)

type Users struct {
	Id                     uint
	Credential_id          uint
	Credential             *Credentials
	UserScheduleCounseling []UserScheduleCounseling `gorm:"foreignKey:User_id;references:Id"`
	UserBooking            []BookingCounseling      `gorm:"foreignKey:User_id;references:Id"`
	First_name             string
	Last_name              string
	Profile_picture        string `gorm:"default:https://pub-86c5755f32914550adb162dd2b8850d0.r2.dev/default-profile.jpg"`
	Phone_number           string
	Birthday               *time.Time
	Status                 string `gorm:"default:INACTIVE"`
	Secret_Otp             *string
	Otp_enable             bool
	Counselor_Favorite     []Counselors `gorm:"many2many:counselor_favorite;foreignKey:Id;references:Id;"`
	CreatedAt              time.Time
}
