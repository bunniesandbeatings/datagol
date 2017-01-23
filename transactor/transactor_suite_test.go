package transactor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTransactor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transactor Suite")
}
