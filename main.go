package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	DEFAULT_HOST_AND_PORT = "0.0.0.0:3000"
)

func writeJSON(w http.ResponseWriter, value any) error {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	return encoder.Encode(value)	
}

func writeError(w http.ResponseWriter, err error) error {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("Content-Type", "application/text")
	_, err = w.Write([]byte(err.Error()))
	return err
}

func main() {
	server := http.NewServeMux()
	
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body := struct {
			Hello string
		} {
			"World",
		}
		if err := writeJSON(w, body); err != nil {
			err = writeError(w, err)
			if err != nil {
				log.Panic(err)
			}
		}
	})

	log.Printf("listening on %s", DEFAULT_HOST_AND_PORT)
	if err := http.ListenAndServe(DEFAULT_HOST_AND_PORT, server); err != nil {
		log.Fatal(err)
	}
}
