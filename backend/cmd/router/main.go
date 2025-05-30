package router

import (
	"pogle-form/backend/cmd/repos"
	authRouter "pogle-form/backend/cmd/router/auth"
)

const (
	PREFIX = "/api"
	GET    = "GET " + PREFIX
	POST   = "POST " + PREFIX
	PUT    = "PUT " + PREFIX
	DELETE = "DELETE " + PREFIX

	ADMIN_VIEW_DIR = "static/dist"
	FORM_VIEW_DIR  = "static/course_register.html"
)

type router struct {
	app *repos.App
}

func Run(
	app *repos.App,
) {

	r := router{
		app: app,
	}
	authRouter.Run(app)
	app.Srv.HandleFunc("GET /", r.handleViewForm)

	app.Srv.Handle(POST+"/courses", r.app.Auth.Authorize(r.handleCreateCourse))
	app.Srv.Handle(GET+"/courses", r.app.Auth.Authorize(r.handleGetCourses))
	app.Srv.HandleFunc(POST+"/persons", r.handleCreatePerson)
	app.Srv.HandleFunc(GET+"/persons", r.handleGetPersons)
}
