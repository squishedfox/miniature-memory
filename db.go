package main

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
)

var (
	forms = make([]*FormModel, 0, 10)
)

func AddForm(newValue CreateFormModel) (string, error) {
	// this is extremely inefficient to look through this swhole collection of records and find the answer
	// but it works for now
	for _, form := range forms {
		if form.Name == newValue.Name {
			return "", fmt.Errorf("Form with name %s already exists", newValue.Name)
		}
	}
	id := uuid.New()
	newFormValue := FormModel{
		ID:          id,
		Name:        newValue.Name,
		CreatedDate: newValue.CreatedDate,
	}
	forms = append(forms, &newFormValue)
	return id.String(), nil
}

func RemoveForm(id string) error {
	forms = slices.DeleteFunc(forms, func(form *FormModel) bool {
		return form.ID.String() == id
	})
	return nil
}

func UpdateForm(id string, updatedValue FormModel) error {
	form, err := GetForm(id)
	if err != nil {
		return err
	}
	if form == nil {
		return fmt.Errorf("No form with the ID %s exists", id)
	}
	form.Name = updatedValue.Name
	return nil // leave for now and leave open for returning an error
}

func GetForm(id string) (*FormModel, error) {
	for _, form := range forms {
		if form.ID.String() == id {
			return form, nil
		}
	}
	return nil, fmt.Errorf("No form with the ID %s exists", id)
}

func GetForms() ([]*FormModel, error) {
	return forms, nil
}
