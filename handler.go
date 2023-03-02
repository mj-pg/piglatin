package main

import (
	"encoding/json"
	"fmt"
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
		log.Printf("parse request error: %v", err)
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
		log.Printf("saving error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		// do not return, translation itself does not error
		// just let client know something went wrong
		// we can still give the output
	}
	log.Println("Output:", translated)

	// format json response
	// {"pig_latin": "some translation value"}
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

	// respond
	if _, err := w.Write(respJson); err != nil {
		log.Printf("response failed: %v", err)
	}
}

// GET /piglatins
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {

	list, err := h.svc.List()
	if err != nil {
		log.Printf("fetching list error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("Output:", list)

	// format json response
	// [
	//  {"text1": "piglatin1"},
	//  {"text2": "piglatin2"},
	//  {"text3": "piglatin3"}
	// ]
	// it may be argued that this type should have been used in the entire app
	// and eliminate the next lines of code just for format processing
	// however [2]string is enough in the other parts(json tags not needed)
	//	and this type seem helpful here only
	type translation struct {
		Text     string `json:"text"`
		PigLatin string `json:"pig_latin"`
	}
	translations := make([]translation, len(list))
	for i, l := range list {
		t := translation{
			Text:     l[0],
			PigLatin: l[1],
		}
		translations[i] = t
	}
	respJson, err := json.Marshal(translations)
	_ = respJson
	// this shouldn't happen
	if err != nil {
		// send output without proper json format
		fmt.Fprintf(w, "%+v", translations)
		return
	}

	// respond
	if _, err := w.Write(respJson); err != nil {
		log.Printf("response failed: %v", err)
	}
}
