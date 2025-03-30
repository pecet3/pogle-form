package dtos

import (
	"encoding/json"
	"io"
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
