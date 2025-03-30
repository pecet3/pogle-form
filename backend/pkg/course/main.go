package course

import "pogle-form/backend/data"

type Course struct {
	d *data.Queries
}

func New(d *data.Queries) *Course {
	return &Course{
		d: d,
	}
}
