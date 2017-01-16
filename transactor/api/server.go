package api

import (
	"github.com/tedsuo/rata"
	"net/http"
	"github.com/bunniesandbeatings/datagol/transactor"
)

func Start(engine *transactor.Connection) error {
	router, err := rata.NewRouter(Routes, NewHandlers(engine))
	if err != nil {
		return err
	}

	err = http.ListenAndServe(":8000", router)
	if err != nil {
		return err
	}

	return nil
}
