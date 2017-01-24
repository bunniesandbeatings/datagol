package testutil_test

import (
	"database/sql"
	"fmt"
	. "github.com/MakeNowJust/heredoc/dot"
	"github.com/olekukonko/tablewriter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"os/exec"
	"time"
)

type DB struct {
	User string
	Name string
	DB   *sql.DB
}

func NewPostgres() *DB {
	db := &DB{
		User: "datagol_test",
		Name: "datagol_test",
	}

	db.recreateDB()
	db.connectDB()

	return db
}

func (db *DB) Close() {
	db.DB.Close()
}

func (db *DB) PrintQuery(query string) {

	rows, err := db.DB.Query(query)

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	table := tablewriter.NewWriter(GinkgoWriter)
	table.SetHeader(columns)

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		sValues := []string{}

		for _, value := range values {
			switch value.(type) {
			case nil:
				sValues = append(sValues, "__NULL__")
			case []byte:
				sValues = append(sValues, string(value.([]byte)))
			default:
				sValues = append(sValues, fmt.Sprint(value))
			}
		}
		table.Append(sValues)
	}
	table.Render()
}

func (db *DB) SingleQuery(query string) interface{} {
	var result interface{}
	err := db.DB.QueryRow(query).Scan(&result)
	if err != nil {
		Fail(fmt.Sprintf("SingleQuery '%s' failed with '%s'", query, err))
	}

	return result
}

func (db *DB) recreateDB() {
	fmt.Fprintf(GinkgoWriter, "** Recreating test database '%s'\n", db.Name)

	script := fmt.Sprintf(
		D(`
	  	psql -c "DROP DATABASE IF EXISTS %s"
		  psql -c "DROP USER IF EXISTS %s"
		  createuser -d %s
		  createdb -E utf8 -e -w -O %s %s
	  `),
		db.Name,
		db.User,
		db.User,
		db.User,
		db.Name)

	cleanDB := exec.Command("bash", "-c", script)

	session, err := gexec.Start(cleanDB, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())

	Eventually(session, time.Second * 10).Should(gexec.Exit(0))
}

func (db *DB) connectDB() {
	fmt.Fprintf(GinkgoWriter, "** Connecting to test database '%s'\n", db.Name)

	var err error
	db.DB, err = sql.Open("postgres", db.datasourceName())
	Expect(err).ToNot(HaveOccurred())

	err = db.DB.Ping()
	Expect(err).ToNot(HaveOccurred())
}

func (db *DB) datasourceName() string {
	return fmt.Sprintf("user=%s dbname=%s sslmode=disable", db.User, db.Name)
}
