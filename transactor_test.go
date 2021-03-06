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
	"bytes"
	"fmt"
	"github.com/bunniesandbeatings/datagol/testutil"
)

var _ = Describe("Transactor", func() {
	var (
		transactor       *Cmd
		session          *Session
		host             = "127.0.0.1"
		port             = "3030"
		apiEndpoint      = fmt.Sprintf("http://%s:%s", host, port)
		entitiesEndpoint = fmt.Sprintf("%s/entities", apiEndpoint)
		testDB           *testutil_test.DB
	)

	BeforeEach(func() {
		testDB = testutil_test.NewPostgres()

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
		Eventually(session).Should(Exit())

		testDB.Close()
	})

	Describe("Assert with lots of entities attributes", func() {
		var response *http.Response
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

		BeforeEach(func() {
			var err error
			response, err = http.Post(entitiesEndpoint, "application/json", bytes.NewBufferString(postData))
			Expect(err).To(BeNil())
		})

		//It("Responds politely", func() {
		//	Expect(response.StatusCode).To(Equal(201))
		//
		//	byteResponse, err := ioutil.ReadAll(response.Body)
		//	Expect(err).To(BeNil())
		//
		//	var jsonResponse = []struct{
		//		StringEntity string `json:"entity"`
		//	}{}
		//
		//	json.Unmarshal(byteResponse, jsonResponse)
		//
		//	Expect(jsonResponse[0]).To(Equal(1))
		//	Expect(jsonResponse[1]).To(Equal(2))
		//})

		It("creates the correct records in the DB", func() {
			Expect(testDB.SingleQuery("SELECT count(DISTINCT time) FROM eavt;")).
				To(Equal(int64(2)))

			Expect(testDB.SingleQuery("SELECT count(DISTINCT entity) FROM eavt;")).
				To(Equal(int64(2)))

			Expect(testDB.SingleQuery("SELECT json_value FROM eavt where attribute='vulnerability/usn';")).
				To(Equal(`"USN-1111-1111"`))

			Expect(testDB.SingleQuery("SELECT json_value FROM eavt where attribute='vulnerability/cvss/base/score';")).
				To(Equal(`4.7`))
		})
	})

})
