package data

import (
	"context"
	"pogle-form/backend/data/dtos"
)

func (c Course) ToDto(d *Queries) *dtos.Course {
	rp, _ := d.GetNumberOfPersonsInCourse(context.Background(), c.ID)
	return &dtos.Course{
		ID:               int(c.ID),
		Name:             c.Name,
		MaxPersons:       int(c.MaxPersons),
		RegisterdPersons: int(rp),
	}
}
