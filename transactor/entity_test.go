package transactor_test

import (
	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
	. "github.com/bunniesandbeatings/datagol/testutil"
)

var _ = FDescribe("Entity", func() {
	var testDB *DB

	BeforeEach(func() {
		testDB = NewPostgres("user=datagol_test dbname=datagol_test sslmode=disable")
	})

	AfterEach(func(){
		testDB.Close()
	})

	Describe("Creating an entity", func() {
		
	})
})
