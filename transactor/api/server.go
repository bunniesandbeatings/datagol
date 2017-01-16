package api

import (
	"github.com/tedsuo/rata"
	"net/http"
	"github.com/bunniesandbeatings/datagol/transactor"
	"os"
	"fmt"
	"log"
)

func Start(engine *transactor.Connection) error {
	router, err := rata.NewRouter(Routes, NewHandlers(engine))
	if err != nil {
		return err
	}

	port := os.Getenv("PORT")
	if port == "" { port = "3000"}

	bind := os.Getenv("BIND")

	listen := fmt.Sprintf("%s:%s", bind, port)

	fmt.Fprintln(os.Stderr, "*********************************")
	fmt.Fprintln(os.Stderr, "* Datagol Transactor starting")
	fmt.Fprintln(os.Stderr, "*")
	fmt.Fprintf(os.Stderr, "* Listening on %s\n", listen)
	fmt.Fprintln(os.Stderr, "* Cltr-C to hang up")
	fmt.Fprintln(os.Stderr, "")
	log.Printf("Listening on %s\n", listen)

	err = http.ListenAndServe(listen, router)
	if err != nil {
		return err
	}

	return nil
}
