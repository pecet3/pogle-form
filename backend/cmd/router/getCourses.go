package router

import (
	"net/http"
	"pogle-form/backend/data/dtos"

	"github.com/pecet3/logger"
)

func (r router) handleGetCourses(w http.ResponseWriter, req *http.Request) {
	courses, err := r.app.Data.ListCourses(req.Context())
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	var dto []*dtos.Course
	for _, c := range courses {
		dto = append(dto, c.ToDto(r.app.Data))
	}
	if err := r.app.Dto.Send(w, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}
