package router

import (
	"database/sql"
	"net/http"
	"pogle-form/backend/data"
	"pogle-form/backend/data/dtos"
	"time"

	"github.com/pecet3/logger"
)

func (r router) handleCreatePerson(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.CreatePerson{}
	if err := r.app.Dto.Get(req, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	p, err := r.app.Data.CreatePerson(req.Context(),
		data.CreatePersonParams{
			Email:    sql.NullString{String: dto.Email, Valid: true},
			FullName: dto.FullName,
			CourseID: int64(dto.ChosenCourseID),
		})
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	c, err := r.app.Data.GetCourse(req.Context(), int64(dto.ChosenCourseID))
	if err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	for _, rc := range dto.ReservedCourseIDs {
		_, err := r.app.Data.CreateChosenReservedCourse(req.Context(),
			data.CreateChosenReservedCourseParams{PersonID: p.ID, CourseID: int64(rc)})
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "course",
			Value:   c.Name,
			Expires: time.Now().Add(time.Hour * 999),
			Path:    "/",
		})
	logger.Debug(dto)
}
