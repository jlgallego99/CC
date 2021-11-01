package obra_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestObra(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Obra Suite")
}
