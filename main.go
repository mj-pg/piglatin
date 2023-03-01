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

	handler := &PigLatinHandler{}
	http.Handle("/piglatins", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Printf("Server exited with error: %v", err)
	}
}

type PigLatinHandler struct{}

// ServeHTTP routes the request based on the HTTP method.
func (h *PigLatinHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Translate(w, r)
		w.WriteHeader(200)
		return
	case http.MethodGet:
		h.List(w, r)
		w.WriteHeader(200)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *PigLatinHandler) Translate(w http.ResponseWriter, r *http.Request) {
}

func (h *PigLatinHandler) List(w http.ResponseWriter, r *http.Request) {
}
