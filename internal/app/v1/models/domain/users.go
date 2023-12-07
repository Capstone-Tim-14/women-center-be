package domain

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id                     uint
	Credential_id          uint
	Credential             *Credentials
	UserScheduleCounseling []UserScheduleCounseling
	UserBooking            []BookingCounseling
	First_name             string
	Last_name              string
	Profile_picture        string `gorm:"default:https://pub-86c5755f32914550adb162dd2b8850d0.r2.dev/default-profile.jpg"`
	Phone_number           string
	Birthday               *time.Time
	Status                 string `gorm:"default:INACTIVE"`
	Secret_Otp             *string
	Otp_enable             bool
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              gorm.DeletedAt
}
