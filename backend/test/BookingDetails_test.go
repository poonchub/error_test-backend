package test

import (
	"test-error/backend/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestQuantity(t *testing.T) {

	govalidator.CustomTypeTagMap.Set("greaterzero", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		num, ok := i.(int)
		if !ok {
			return false
		}
		return num > 0
	}))

	g := NewGomegaWithT(t)

	t.Run(`quantity is required`, func(t *testing.T) {
		bkd := entity.BookingDetails{
			Quantity: 0,	// ผิดตรงนี้
			BookingID: 1,
		}

		ok, err := govalidator.ValidateStruct(bkd)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(ContainSubstring("Quantity is required"))
	})

	t.Run(`quantity must be greater than 0`, func(t *testing.T) {
		bkd := entity.BookingDetails{
			Quantity: -5,	// ผิดตรงนี้
			BookingID: 1,
		}

		ok, err := govalidator.ValidateStruct(bkd)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(ContainSubstring("Quantity must be greater than 0"))

		// g.Expect(err.Error()).To(Equal("Quantity must be greater than 0"))
		
		/*
			ใช้ ContainSubstring() เพราะค่า error ที่ได้มี error อันอื่นติดมาด้วย

			=== RUN   TestQuantity/quantity_must_be_greater_than_0
			Expected
            	<string>: Quantity must be greater than 0;TotalPrice is required
        	to equal
            	<string>: Quantity must be greater than 0
		*/
	})

	t.Run(`quantity is valid`, func(t *testing.T) {
		bkd := entity.BookingDetails{
			Quantity: 5,
			BookingID: 1,
		}

		ok, err := govalidator.ValidateStruct(bkd)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

		// g.Expect(ok).NotTo(BeTrue())
		// g.Expect(err).NotTo(BeNil())

		// g.Expect(err.Error()).To(ContainSubstring("gdfdhdh"))

		/*
			ต้องการตรวจสอบให้ถูกต้อง แต่ติด error ของ field Totalprice ของตาราง Bookings

			=== RUN   TestQuantity/quantity_is_valid
			Expected
            	<string>: TotalPrice is required
        	to contain substring
            	<string>: gdfdhdh
		*/
	})
}