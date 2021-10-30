package cancion_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCancion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cancion Suite")
}
