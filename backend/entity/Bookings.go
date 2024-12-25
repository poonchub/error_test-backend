package entity

import (
	"time"

	"gorm.io/gorm"
)

type Bookings struct {
	gorm.Model
	BookingDate		time.Time
	TotalPrice		float32	`valid:"required~TotalPrice is required,greaterzero~TotalPrice must be greater than 0"`
	TotalQuantity	int
	SpecialRequest	string

	BookingDetails	[]BookingDetails	`gorm:"foreignKey:BookingID"`
}