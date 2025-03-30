package router

import (
	"net/http"
	"pogle-form/backend/data/dtos"

	"github.com/pecet3/logger"
)

func (r router) handleCreatePerson(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.CreatePerson{}
	if err := r.app.Dto.Get(req, dto); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logger.Debug(dto)
}
