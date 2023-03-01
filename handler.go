package main

import "net/http"

type Handler struct{}

// ServeHTTP routes the request based on the HTTP method.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Translate(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
}
