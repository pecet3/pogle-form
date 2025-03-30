package dtos

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type CreatePerson struct {
	FullName          string `json:"full_name" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	ChosenCourseID    int    `json:"chosen_course" validate:"required"`
	ReservedCourseIDs []int  `json:"reserved_courses"`
}

func (c CreatePerson) Validate(v *validator.Validate) error {
	err := v.Struct(c)
	if err != nil {
		return err
	}
	return nil
}

type Person struct {
	ID           int     `json:"id"`
	FullName     string  `json:"full_name"`
	Email        string  `json:"email" `
	ChosenCourse *Course `json:"chosen_course" `
}

func (p Person) Send(w io.Writer) error {
	if err := json.NewEncoder(w).Encode(&p); err != nil {
		return err
	}
	return nil
}
