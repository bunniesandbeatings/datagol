package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
)

type Options struct {
	LogTarget string `long:"log-file" description:"Where to log to. Defaults to stderr" env:"TRANSACTOR_LOG"`
}

var globalOptions Options

var parser = flags.NewParser(&globalOptions, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			log.Printf("Exited with error: %s", err)
			os.Exit(1)
		}
	}
}

func ConfigureLogging() {
	if globalOptions.LogTarget == "" {
		log.SetOutput(os.Stderr)
	} else {
		logFile, err := os.Open(globalOptions.LogTarget)
		if err != nil {
			panic(fmt.Sprintf("Cannot open log target '%s' for append.", globalOptions.LogTarget))
		}
		defer logFile.Close()
		log.SetOutput(logFile)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
