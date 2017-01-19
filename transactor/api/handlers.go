package api

import (
	"github.com/tedsuo/rata"
	"github.com/bunniesandbeatings/datagol/transactor/api/handlers"
	"github.com/bunniesandbeatings/datagol/transactor"
)

func NewHandlers(connection *transactor.Connection) rata.Handlers {
	return rata.Handlers{
		Accumulate: handlers.NewAccumulateHandler(connection),
		Assert: handlers.NewAssertHandler(connection),
		Docs: handlers.NewDocsHandler(),
	}
}
