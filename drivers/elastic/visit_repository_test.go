package elastic

import (
	"testing"

	"github.com/index0h/go-tracker/components"
)

func Test_VisitRepository_Interface(t *testing.T) {
	func(event components.VisitRepositoryInterface) {}(&VisitRepository{})
}
