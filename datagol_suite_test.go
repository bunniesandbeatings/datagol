package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	"testing"
)

var datagolCLI string

func TestDatagol(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Datagol Suite")
}

var _ = BeforeSuite(func() {
	var err error
	datagolCLI, err = Build("github.com/bunniesandbeatings/datagol")
	Expect(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	CleanupBuildArtifacts()
})