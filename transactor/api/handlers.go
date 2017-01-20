package api

import (
	"github.com/tedsuo/rata"
	"github.com/bunniesandbeatings/datagol/transactor/api/handlers"
	"github.com/bunniesandbeatings/datagol/transactor"
)

func NewHandlers(connection *transactor.Connection) rata.Handlers {
	return rata.Handlers{
		Assert: handlers.NewAssertHandler(connection),
		Accumulate: handlers.NewAccumulateHandler(connection),
		Docs: handlers.NewDocsHandler(),
	}
}
