package dtos

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type Course struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	MaxPersons       int    `json:"max_persons"`
	RegisterdPersons int    `json:"registered_persons"`
}

func (r Course) Send(w io.Writer) error {
	if err := json.NewEncoder(w).Encode(&r); err != nil {
		return err
	}
	return nil
}

type CreateCourse struct {
	Name       string `json:"name" validate:"required"`
	MaxPersons int    `json:"max_persons" validate:"required"`
}

func (c CreateCourse) Validate(v *validator.Validate) error {
	err := v.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
