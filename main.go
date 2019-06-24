package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tty2/company/controller"
	"github.com/tty2/company/datastore"
	"github.com/tty2/company/handler"
)

const port string = ":8080"

func main() {
	config := "config string"
	ds := datastore.New(config)
	service := controller.New(ds.Employee)
	api := handler.New(service)
	fmt.Printf("running service on port %s\n", port)
	log.Panic(http.ListenAndServe(port, api.Mux))
}
