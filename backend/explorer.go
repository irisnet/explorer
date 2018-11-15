package main

import (
	"github.com/irisnet/explorer/backend/modules"
	"github.com/irisnet/explorer/backend/task"
	"log"
)

func main() {
	go task.Start()
	server := modules.NewServer()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe Failed: ", err)
	}
}
