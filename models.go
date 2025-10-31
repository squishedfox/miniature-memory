package main

type (
	FormModel struct {
		ID string `json:"id"`
		Name string `json:"name"`
		CreatedDate string `json:"createdDate"`
	}
	GetFormsListResponse struct {
		Forms []FormModel `json:"forms"`
		Count int64 `json:"count"`
	}
)
