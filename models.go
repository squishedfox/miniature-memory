package main

import "github.com/google/uuid"

type (
	CreateFormModel struct {
		// Name is the human readable name of the form
		Name string `json:"name"`
		// CreatedDate is the unix epoch since January 1st, 1970
		CreatedDate int64 `json:"createdDate"`
	}
	FormModel struct {
		// ID for tracking the unique identifier of the ID
		ID uuid.UUID `json:"id"`
		// Name is the human readable name of the form
		Name string `json:"name"`
		// CreatedDate is the unix epoch since January 1st, 1970
		CreatedDate int64 `json:"createdDate"`
	}
	GetFormsListResponse struct {
		Forms []FormModel `json:"forms"`
		Count int64 `json:"count"`
	}
)
