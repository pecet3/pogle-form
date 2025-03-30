package router

import "pogle-form/backend/cmd/repos"

const (
	PREFIX = "/api"
	GET    = "GET " + PREFIX
	POST   = "POST " + PREFIX
	PUT    = "PUT " + PREFIX
	DELETE = "DELETE " + PREFIX

	VIEW_DIR = "./cmd/view"
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
	// authRouter.Run(app)

	app.Srv.HandleFunc("/", r.handleView)

}
