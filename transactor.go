package main

import (
	"os"
	"github.com/bunniesandbeatings/datagol/transactor"
	"log"
	"github.com/bunniesandbeatings/datagol/transactor/api"
	"fmt"
)

type APIOptions struct {
	Address string `short:"a" long:"address" description:"interface to bind to" env:"ADDRESS" default:"0.0.0.0"`
	Port    string `short:"p" long:"port" description:"port to bind to" env:"PORT" default:"3000"`
}

type DBOptions struct {
	DBConnectString string `short:"d" long:"datasource" description:"Postgresql datasource string" env:"DB_CONNECT_STRING" default:"user=datagol dbname=datagol sslmode=disable"`
}

type TransactorOptions struct {
	API *APIOptions `group:"API Server"`
	DB  *DBOptions  `group:"Backing database"`
}

var transactorOptions TransactorOptions

func (transactorOptions *TransactorOptions) Execute(args []string) error {

	ConfigureLogging()

	connection, err := transactor.NewConnection(transactorOptions.DB.DBConnectString)
	if err != nil {
		return err
	}

	address := fmt.Sprintf("%s:%s", transactorOptions.API.Address, transactorOptions.API.Port)

	server := api.ServerConfig{ Address: address }

	err = server.Start(connection)
	if err != nil {
		log.Printf("Server terminated with: %s", err)
		return nil
	}

	return nil
}

func init() {
	transactorCommand, err := parser.AddCommand(
		"start-transactor",
		"Start the transactor server",
		"The transactor runs a server bound to host:port backed by postgres",
		&transactorOptions)

	if err != nil {
		panic(err)
		os.Exit(1)
	}

	transactorCommand.Aliases = []string{"transactor"}
}
