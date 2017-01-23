package testutil_test

import (
	"fmt"
	"database/sql"
	"time"
	"os/exec"
	"github.com/onsi/gomega/gexec"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


type DB struct {
	DatasourceName string
	DB *sql.DB
}

func NewPostgres(datasourceName string) *DB {
	db := &DB{
		DatasourceName: datasourceName,
	}

	cleanTestDB()

	var err error
	db.DB, err = sql.Open("postgres", db.DatasourceName)
	Expect(err).ToNot(HaveOccurred())

	err = db.DB.Ping()
	Expect(err).ToNot(HaveOccurred())

	return db
}

func (db *DB) Close() {
	db.DB.Close()
}

func cleanTestDB() {
	cleanDB := exec.Command("./script/ensure-test-db")
	session, err := gexec.Start(cleanDB, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Eventually(session).Should(gexec.Exit(0))
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

func (db *DB) GetRows(query string) []Row {
	queryRows, err := db.DB.Query(query)
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

func (db *DB) PrintTable() {
	rows := db.GetRows("SELECT * FROM eavt;")

	for _, row := range rows {
		fmt.Println(row.String())
	}
}
