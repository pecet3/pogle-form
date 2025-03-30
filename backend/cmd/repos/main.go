package repos

import (
	"net/http"
	"pogle-form/backend/data"
	"pogle-form/backend/data/dtos"
	"pogle-form/backend/pkg/auth"
	"pogle-form/backend/pkg/course"
)

type App struct {
	Srv    *http.ServeMux
	Data   *data.Queries
	Auth   *auth.Auth
	Dto    dtos.Dto
	Course course.Course
}

func NewApp() *App {
	mux := http.NewServeMux()
	db := data.NewSQLite()
	data := data.New(db)
	auth := auth.New()
	return &App{
		Srv:    mux,
		Data:   data,
		Dto:    dtos.New(),
		Auth:   auth,
		Course: *course.New(data),
	}

}
