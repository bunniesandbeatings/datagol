package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
	. "os/exec"

	"testing"
	"database/sql"
	"time"
	"fmt"
)

var testDatasourceName = "user=datagol_test dbname=datagol_test sslmode=disable"
var datagolCLI string
var testDB *sql.DB

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

func CleanTestDB() {
	cleanDB := Command("./script/ensure-test-db")
	session, err := Start(cleanDB, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(Exit(0))
}

func ResetTestDB() {
	if testDB != nil {
		if err := testDB.Ping(); err == nil {
			testDB.Close()
		}
		testDB = nil
	}

	CleanTestDB()

	var err error
	testDB, err = sql.Open("postgres", testDatasourceName)
	Expect(err).ToNot(HaveOccurred())

	err = testDB.Ping()
	Expect(err).ToNot(HaveOccurred())
}

type Row struct {
	entity    uint64
	attribute string
	jsonValue string
	time      time.Time
}

func (row *Row) String() string {
	return fmt.Sprintf("[%d@%s] %s:%s", row.entity, row.time.Local(), row.attribute, row.jsonValue)
}

func GetRows(query string) []Row {
	queryRows, err := testDB.Query(query)
	Expect(err).To(BeNil())
	defer queryRows.Close()

	rows := []Row{}

	for queryRows.Next() {
		row := Row{}
		queryRows.Scan(&row.entity, &row.attribute, &row.jsonValue, &row.time)
		rows = append(rows, row)
	}

	return rows
}


func PrintTable() {
	rows := GetRows("SELECT * FROM eavt;")

	for _, row := range rows {
		fmt.Println(row.String())
	}
}