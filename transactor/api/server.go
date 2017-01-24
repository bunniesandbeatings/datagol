package api

import (
	"github.com/tedsuo/rata"
	"net/http"
	"github.com/bunniesandbeatings/datagol/transactor"
	"os"
	"fmt"
	"log"
)

type ServerConfig struct {
	Address string
}

func (serverConfig *ServerConfig) Start(engine *transactor.Backend) error {
	router, err := rata.NewRouter(Routes, NewHandlers(engine))
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "*********************************")
	fmt.Fprintln(os.Stderr, "* Datagol Transactor starting")
	fmt.Fprintln(os.Stderr, "*")
	fmt.Fprintf(os.Stderr, "* Listening on %s\n", serverConfig.Address)
	fmt.Fprintln(os.Stderr, "* Cltr-C to hang up")
	fmt.Fprintln(os.Stderr, "")
	log.Printf("Listening on %s\n", serverConfig.Address)

	err = http.ListenAndServe(serverConfig.Address, router)
	if err != nil {
		return err
	}

	return nil
}
