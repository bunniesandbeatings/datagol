package handlers

import (
	"net/http"
	"github.com/bunniesandbeatings/datagol/transactor"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)

type AssertHandler struct {
	Connection *transactor.Connection
}

func NewAssertHandler(connection *transactor.Connection) *AssertHandler {
	return &AssertHandler{
		Connection: connection,
	}
}

func (handler *AssertHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()

	var entity = transactor.AttributeValuesJson{}

	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal(requestBody, &entity)


	entityId, err := handler.Connection.CreateEntity(entity)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	result := fmt.Sprintf("Created Entity: %d\n", entityId)

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprint(writer, result)
}
