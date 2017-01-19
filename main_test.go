package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	. "github.com/onsi/gomega/gbytes"
	. "os/exec"
	. "github.com/MakeNowJust/heredoc/dot"
	"regexp"
	"net/http"
	"fmt"
	"bytes"
	"io/ioutil"
)

var _ = Describe("Transactor", func() {
	var (
		transactor  *Cmd
		session     *Session
		host        = "127.0.0.1"
		port        = "3030"
		apiEndpoint = fmt.Sprintf("http://%s:%s", host, port)
		entitiesEndpoint = fmt.Sprintf("%s/entities", apiEndpoint)
	)

	BeforeEach(func() {
		CleanTestDB()

		transactor = Command(
			datagolCLI,
			"start-transactor",
			"-a", host,
			"-p", port,
			"-d", "user=datagol_test dbname=datagol_test sslmode=disable",
		)

		var err error
		session, err = Start(transactor, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
		Eventually(session.Err).Should(Say(regexp.QuoteMeta(`Listening on 127.0.0.1:3030`)))
	})

	AfterEach(func() {
		session.Interrupt()
		//Eventually(session.Out).Should(Say(`Got signal`))
		Eventually(session).Should(Exit())
	})

	Describe("Assert with lots of entities and attributes", func() {
		// todo db connection fails
		var postData = D(`[
		  {
				"testing/integers": 1,
				"vulnerability/cve/id": "CVE-1111-1111",
				"vulnerability/cvss/base/score": 4.7,
				"vulnerability/cvss/base/severity": "medium",
				"vulnerability/cvss/base/vector-string": "CVSS:3.0/AV:A/AC:H/PR:L/UI:R/S:U/C:H/I:N/A:N"
			},
			{
				"vulnerability/usn": "USN-1111-1111",
				"vulnerability/cves": ["CVE-1111-1111", "CVE-1111-1112"]
			}
		]`)

		It("Responds like a rest endpoint should", func() {
			response, err := http.Post(entitiesEndpoint, "application/json", bytes.NewBufferString(postData))
			Expect(err).To(BeNil())

			Expect(response.StatusCode).To(Equal(201))

			byteResponse, _ := ioutil.ReadAll(response.Body)
			stringResponse := string(byteResponse)
			Expect(stringResponse).To(ContainSubstring(`Created Entity: 1`))

		})

		It("creates the correct records in the DB", func() {

		})

	})
})
