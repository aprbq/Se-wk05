package main

import (
	"se67-wk05/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestMember(t *testing.T) {

	g := NewGomegaWithT(t)
	t.Run(`all valid`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "0123456789",
			Password:    "1234",
			Url:         "https://www.linkedin.com/company/ilink/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})

	t.Run("phone_number is invalid", func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "123456780014", // ผิดตรงนี้ (ไม่ครบ 10 หลัก)
			Password:    "1111",
			Url:         "https://www.example.com",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone Number length is not 10 digits."))
	})

	t.Run("URL is invalid", func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "1234567890",
			Password:    "1234",
			Url:         "dsfdssdf", // ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Url is invalid"))
	})
}

