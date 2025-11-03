package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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
			http.NotFound(w, r)
			return
		}
		form, err := GetForm(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if form == nil {
			http.NotFound(w, r)
			return
		}
		err = writeJSON(w, form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	CreateFormHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		form := CreateFormModel{
			Name:        r.PostFormValue("name"),
			CreatedDate: time.Now().UTC().Unix(),
		}
		id, err := AddForm(form)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Location", fmt.Sprintf("/forms/%s", id))
		w.WriteHeader(http.StatusCreated)
	})

	DeleteFormHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if len(id) == 0 {
			http.NotFound(w, r)
			return
		}
		err := RemoveForm(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	UpdateFormHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if len(id) == 0 {
			http.NotFound(w, r)
			return
		}
		formModel := FormModel{}
		reader := json.NewDecoder(r.Body)
		reader.Decode(&formModel)
		err := UpdateForm(id, formModel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
)
