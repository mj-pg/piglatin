package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	svc Service
}

// ServeHTTP routes the request based on the HTTP method.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Translate(w, r)
	case http.MethodGet:
		h.List(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// POST /piglatins
func (h *Handler) Translate(w http.ResponseWriter, r *http.Request) {

	// parse request
	//
	reqBody := struct {
		Text string `json:"text"`
	}{}

	d := json.NewDecoder(r.Body)
	if err := d.Decode(&reqBody); err != nil {
		log.Println("parse request error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Got input: %+v", reqBody)

	if reqBody.Text == "" {
		log.Println("Empty text")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// translate input to pig latin
	//
	translated, err := h.svc.Translate(reqBody.Text)
	if err != nil {
		log.Println("saving error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		// do not return, translation itself does not error
		// we can still give the output
	}
	log.Println("Output:", translated)

	// respond translation
	//
	respBody := struct {
		PigLatin string `json:"pig_latin"`
	}{
		PigLatin: translated,
	}
	respJson, err := json.Marshal(respBody)
	// this shouldn't happen
	if err != nil {
		// send raw text
		w.Write([]byte(translated))
		return
	}
	w.Write(respJson)
}

// GET /piglatins
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {

}
