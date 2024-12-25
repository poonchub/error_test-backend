package test

import (
	"test-error/backend/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestTotalPrice(t *testing.T) {

	govalidator.CustomTypeTagMap.Set("greaterzero", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		num, ok := i.(float32)
		if !ok {
			return false
		}
		return num > 0
	}))

	g := NewGomegaWithT(t)

	t.Run(`total_price is required`, func(t *testing.T) {
		price := entity.Bookings{
			TotalPrice: 0,	// ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(price)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("TotalPrice is required"))
	})

	t.Run(`total_price must be greater than 0`, func(t *testing.T) {
		price := entity.Bookings{
			TotalPrice: -10000,	// ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(price)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("TotalPrice must be greater than 0"))
	})

	t.Run(`total_price is valid`, func(t *testing.T) {
		price := entity.Bookings{
			TotalPrice: 10000,	// ถูกต้อง
		}

		ok, err := govalidator.ValidateStruct(price)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}