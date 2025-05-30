package router

import (
	"net/http"
	"pogle-form/backend/data"
	"pogle-form/backend/data/dtos"

	"github.com/pecet3/logger"
)

func (r router) handleCreateCourse(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.CreateCourse{}
	if err := r.app.Dto.Get(req, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	c, err := r.app.Data.CreateCourse(req.Context(), data.CreateCourseParams{
		Name:       dto.Name,
		MaxPersons: int64(dto.MaxPersons),
	})
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if err := r.app.Dto.Send(w, c.ToDto(r.app.Data)); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	logger.Info("Created a new course: ", dto)
}
