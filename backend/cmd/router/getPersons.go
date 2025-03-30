package router

import (
	"net/http"
	"pogle-form/backend/data/dtos"

	"github.com/pecet3/logger"
)

func (r router) handleGetPersons(w http.ResponseWriter, req *http.Request) {
	var dto []*dtos.Person
	persons, err := r.app.Data.ListPersons(req.Context())
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	for _, p := range persons {
		dto = append(dto, p.ToDto(r.app.Data))
	}
	if err := r.app.Dto.Send(w, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}
