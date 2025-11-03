package main

import (
	"log"
	"net/http"
)

const (
	DEFAULT_HOST_AND_PORT = "0.0.0.0:3000"
)

func main() {
	server := http.NewServeMux()
	
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			OperationNotSupported(w, r)
			return
		}
		body := struct {
			Hello string
		} {
			"World",
		}
		if err := writeJSON(w, body); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	server.HandleFunc("/forms", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodGet:
				GetFormsListHandler(w, r)
			case http.MethodPost:
				CreateFormHandler(w, r)
			default:
				OperationNotSupported(w, r)
		}

	})
	server.HandleFunc("/forms/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodGet:
				GetFormHandler(w, r)
			case http.MethodPut:
				UpdateFormHandler(w, r)
			case http.MethodDelete:
				DeleteFormHandler(w, r)
			default:
				OperationNotSupported(w, r)	

		}
	})
	log.Printf("listening on %s", DEFAULT_HOST_AND_PORT)
	if err := http.ListenAndServe(DEFAULT_HOST_AND_PORT, server); err != nil {
		log.Fatal(err)
	}
}
