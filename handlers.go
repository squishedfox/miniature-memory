package main

import "net/http"

var (
	NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	OperationNotSupported = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	GetFormsListHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		forms, err := GetForms()
		err = writeJSON(w, forms)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	})

	GetFormHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")	
		if len(id) == 0 {
			NotFoundHandler(w, r)
			return
		}
		form, err := GetForm(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if form == nil {
			http.NotFoundHandler() 
			return
		}
		err = writeJSON(w, form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	CreateFormHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	DeleteFormHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})

	UpdateFormHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
)
