package entity

import "gorm.io/gorm"

type BookingDetails struct {
	gorm.Model
	Quantity   	int 	`valid:"required~Quantity is required,greaterzero~Quantity must be greater than 0"`

	BookingID	uint
	Booking		Bookings	`gorm:"foreignKey:BookingID"`

}