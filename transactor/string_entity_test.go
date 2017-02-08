package transactor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/bunniesandbeatings/datagol/testutil"
	"github.com/bunniesandbeatings/datagol/transactor"
	"github.com/bunniesandbeatings/datagol/attributes"
	"time"
)

var _ = Describe("StringEntity", func() {
	var (
		testDB       *DB
		backend      *transactor.Backend
		err          error
		stringEntity transactor.StringEntity
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
			stringEntity = transactor.StringEntity{
				"attr/String":  "\"Once Upon\nA Time\"",
				"attr/Number":  "-2.75e+3",
				"attr/Boolean": "true",
				"attr/Object":  "{\"first\": 1, \"second\":[1,2]}",
				"attr/Array":   "[1,2,\"three\",{\"Four\":4}]",
				"attr/Null":    "null",
			}
		})

		It("creates each row with verbatim values", func() {
			_, err := backend.Commit(stringEntity)
			Expect(err).To(BeNil())

			Expect(testDB.SingleQuery("SELECT DISTINCT entity FROM eavt")).
				To(Equal(int64(1)))

			Expect(testDB.SingleQuery("SELECT count(DISTINCT time) FROM eavt")).
				To(Equal(int64(1)))

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

		FIt("returns sequenced IDs and timestamps with entities", func() {
			updatedEntity, _ := backend.Commit(stringEntity)
			Expect(stringEntity["attr/String"]).To(Equal("\"Once Upon\nA Time\""))
			Expect(updatedEntity[attributes.IDENT]).To(Equal("1"))

			assertTime := updatedEntity[attributes.TIME]

			Expect(assertTime).
				To(ContainSubstring(time.Now().Format("2006-01-02T15")))

		})

	})
	Context("Multiple entities at once", func() {
		//BeforeEach(func() {
		//	stringEntities := []transactor.StringEntity{
		//		{"attr/1":  "Entity one"},
		//		{"attr/2":  "Entity two"},
		//		{"attr/3":  "Entity three"},
		//		{"attr/4":  "Entity four"},
		//		{"attr/5":  "Entity five"},
		//		{"attr/6":  "Entity six"},
		//	}
		//})

		//It("returns sequenced IDs and timestamps", func() {
		//	ids, err := backend.Commit(stringEntity)
		//	Expect(err).To(BeNil())
		//
		//
		//})

		Context("Massive concurrency", func() {
			It("Survives", func() {

			})
			It("gets all unique ids", func() {

			})
		})

		Context("placeholders for relations", func() {
			It("Relates two record with placeholders", func() {

			})
		})

		Context("Emtpy array", func() {

		})

		Context("Empty string entity", func() {

		})
	})
})
