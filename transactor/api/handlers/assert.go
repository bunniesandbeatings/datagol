package handlers

import (
	"net/http"
	"github.com/bunniesandbeatings/datagol/transactor"
	"encoding/json"
	"io/ioutil"
)

type AssertHandler struct {
	Connection *transactor.Connection
}

func NewAssertHandler(connection *transactor.Connection) *AssertHandler {
	return &AssertHandler{
		Connection: connection,
	}
}

type Entities []Entity


func (handler *AssertHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var entities = Entities{}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal(requestBody, &entities)

	for i, entity := range entities {
		handler.Connection.CreateEntity(entity)
	}
}
