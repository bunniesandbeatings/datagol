package api

import (
	"net/http"
	"github.com/bunniesandbeatings/datagol/transactor"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	"github.com/bunniesandbeatings/datagol/attributes"
)

type AssertHandler struct {
	Connection *transactor.Backend
}

func NewAssertHandler(connection *transactor.Backend) *AssertHandler {
	return &AssertHandler{
		Connection: connection,
	}
}

type AssertEntities []transactor.StringEntity

func (handler *AssertHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var entities = AssertEntities{}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal(requestBody, &entities)

	result := ""

	for _, entity := range entities {

		newEntity, err := handler.Connection.Commit(entity)
		if err != nil {
			log.Println(err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		result = result + fmt.Sprintf("Created StringEntity: %d\n", newEntity[attributes.IDENT])
	}

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprint(writer, result)
}
