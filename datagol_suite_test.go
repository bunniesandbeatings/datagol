package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	. "os/exec"

	"testing"
)

func TestDatagol(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Datagol Suite")
}

var datagolCLI string

var _ = BeforeSuite(func() {
	var err error
	datagolCLI, err = Build("github.com/bunniesandbeatings/datagol")
	Expect(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	CleanupBuildArtifacts()
})

var CleanTestDB = func() {
	cleanDB := Command("./script/ensure-test-db")
	session, err := Start(cleanDB, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(Exit(0))

}
