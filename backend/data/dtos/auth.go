package dtos

import "github.com/go-playground/validator/v10"

type Login struct {
	Name     string `json:"name" validate:"required,min=2,max=16,alphanumunicode"`
	Password string `json:"email" validate:"required,max=64"`
}

func (l Login) Validate(v *validator.Validate) error {
	err := v.Struct(l)
	if err != nil {
		return err
	}
	return nil
}
