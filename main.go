package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("vim-go")

	// server port
	//
	var port int
	flag.IntVar(&port, "port", 3000, "port for this service")

	// init service
	//
	var store Store
	svc := Service{store}

	// init controller
	//
	h := &Handler{svc}
	http.Handle("/piglatins", h)

	// start server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Printf("Server exited with error: %v", err)
	}
}
