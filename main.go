package main

import (
	_ "github.com/lib/pq"
	"log"
	"os"
	"github.com/bunniesandbeatings/datagol/transactor/api"
	"github.com/bunniesandbeatings/datagol/transactor"
)


func main() {
	log.SetOutput(os.Stderr)

	connection, err := transactor.NewConnection("user=datagol dbname=datagol sslmode=disable")

	if err != nil {
		log.Fatal("Error: Could not build a connection: ", err)
	}

	err = api.Start(connection)
	if err != nil {
		log.Fatal("API Server exited: ", err)
	}

}
