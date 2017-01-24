package transactor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/bunniesandbeatings/datagol/testutil"
	"github.com/bunniesandbeatings/datagol/transactor"
)

var _ = Describe("StringEntity", func() {
	var (
		testDB       *DB
		backend      *transactor.Backend
		stringEntity *transactor.StringEntity
		err          error
	)

	BeforeEach(func() {
		testDB = NewPostgres()

		backend, err = transactor.NewBackend("user=datagol_test dbname=datagol_test sslmode=disable")
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		testDB.Close()
		backend.Shutdown()
	})

	Context("Given an entity with different JS serializations", func() {
		BeforeEach(func() {
			stringEntity = &transactor.StringEntity{
				"attr/String": "\"Once Upon\nA Time\"",
				"attr/Number": "-2.75e+3",
				"attr/Boolean": "true",
				"attr/Object": "{\"first\": 1, \"second\":[1,2]}",
				"attr/Array": "[1,2,\"three\",{\"Four\":4}]",
				"attr/Null": "null",
			}
		})

		It("creates each row with verbatim values", func() {
			id, err := backend.Commit(stringEntity)

			Expect(err).To(BeNil())
			Expect(id).To(Equal(uint64(1)))

			Expect(testDB.SingleQuery("SELECT value FROM eavt WHERE attribute='attr/String'")).
				To(Equal("\"Once Upon\nA Time\""))

			Expect(testDB.SingleQuery("SELECT value FROM eavt WHERE attribute='attr/Number'")).
				To(Equal("-2.75e+3"))

			Expect(testDB.SingleQuery("SELECT value FROM eavt WHERE attribute='attr/Boolean'")).
				To(Equal("true"))

			Expect(testDB.SingleQuery("SELECT value FROM eavt WHERE attribute='attr/Object'")).
				To(Equal("{\"first\": 1, \"second\":[1,2]}"))

			Expect(testDB.SingleQuery("SELECT value FROM eavt WHERE attribute='attr/Array'")).
				To(Equal("[1,2,\"three\",{\"Four\":4}]"))

			Expect(testDB.SingleQuery("SELECT value FROM eavt WHERE attribute='attr/Null'")).
				To(Equal("null"))
		})

	})
})
