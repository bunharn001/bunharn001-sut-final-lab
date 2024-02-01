package unit

import (
	"testing"
	// "fmt"
	"github.com/asaskevich/govalidator"
	"github.com/bunharn001/bunharn001-sut-final-lab/entity"
	. "github.com/onsi/gomega"
)

func test_success(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(``, func(t *testing.T) {
		employee := entity.Employee{
			Name:       "Somchai",
			URL:        "",
			EmployeeID: "EM1234567890",
		}

		ok, err := govalidator.ValidateStruct(employee)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
