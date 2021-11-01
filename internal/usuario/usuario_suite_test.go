package usuario_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUsuario(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usuario Suite")
}
