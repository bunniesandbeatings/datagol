package api

import 	(
	"github.com/tedsuo/rata"
	"github.com/bunniesandbeatings/datagol/transactor"
)

func NewHandlers(connection *transactor.Backend) rata.Handlers {
	return rata.Handlers{
		Assert: NewAssertHandler(connection),
	}
}
