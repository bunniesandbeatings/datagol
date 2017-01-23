package api

import 	(
	"github.com/tedsuo/rata"
	"github.com/bunniesandbeatings/datagol/transactor"
)

func NewHandlers(connection *transactor.Connection) rata.Handlers {
	return rata.Handlers{
		Assert: NewAssertHandler(connection),
	}
}
