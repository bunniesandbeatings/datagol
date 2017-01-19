package handlers

import (
	"net/http"
	"github.com/bunniesandbeatings/datagol/transactor"
)

type AccumulateHandler struct {
	Connection *transactor.Connection
}

func NewAccumulateHandler(connection *transactor.Connection) *AccumulateHandler {
	return &AccumulateHandler{
		Connection: connection,
	}
}

//type AssertEntities []transactor.AttributeValuesJson

func (handler *AccumulateHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	//var entities = AssertEntities{}
	//
	//requestBody, err := ioutil.ReadAll(request.Body)
	//if err != nil {
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//
	//json.Unmarshal(requestBody, &entities)
	//
	//result := ""
	//
	//for _, entity := range entities {
	//
	//	entityId, err := handler.Connection.CreateEntity(entity)
	//	if err != nil {
	//		log.Println(err)
	//		http.Error(writer, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	result = result + fmt.Sprintf("Created Entity: %d\n", entityId)
	//}
	//
	//fmt.Fprint(writer, result)

}