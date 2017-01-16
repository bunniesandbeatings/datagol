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

	log.Println("*********************************")
	log.Println("* Datagol Transactor starting")
	log.Printf("Listening on %s\n", listen)

	err = http.ListenAndServe(listen, router)
	if err != nil {
		return err
	}

	return nil
}
