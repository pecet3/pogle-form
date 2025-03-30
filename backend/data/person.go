package data

import (
	"context"
	"pogle-form/backend/data/dtos"
)

func (p Person) ToDto(d *Queries) *dtos.Person {
	c, _ := d.GetCourse(context.Background(), p.CourseID)
	return &dtos.Person{
		ID:           int(p.ID),
		FullName:     p.FullName,
		Email:        p.Email.String,
		ChosenCourse: c.ToDto(d),
	}
}
