package api

import (
	"github.com/tedsuo/rata"
	"github.com/bunniesandbeatings/datagol/transactor/api/handlers"
	"github.com/bunniesandbeatings/datagol/transactor"
)

func NewHandlers(engine *transactor.Connection) rata.Handlers {
	return rata.Handlers{
		//Accumulate: newAccumulateHandler(),
		Assert: handlers.NewAssertHandler(engine),
		Docs: handlers.NewDocsHandler(),
	}
}
