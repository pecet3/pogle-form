package dtos

import "github.com/go-playground/validator/v10"

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
