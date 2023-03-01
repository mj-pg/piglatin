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

	// init server
	//
	h := &Handler{}
	http.Handle("/piglatins", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Printf("Server exited with error: %v", err)
	}
}
