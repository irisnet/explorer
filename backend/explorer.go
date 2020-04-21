package main

import (
	"github.com/irisnet/explorer/backend/rest"
	"log"
)

func main() {
	//task.Start()
	server := rest.NewApiServer()

	if err := server.Start(); err != nil {
		log.Fatal("ListenAndServe Failed: ", err)
	}
}
