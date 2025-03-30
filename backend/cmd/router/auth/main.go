package authRouter

import "pogle-form/backend/cmd/repos"

type router struct {
	app *repos.App
}

const (
	PREFIX = "/api/auth"
	GET    = "GET " + PREFIX
	POST   = "POST " + PREFIX
	PUT    = "PUT " + PREFIX
	DELETE = "DELETE " + PREFIX
)

func Run(
	app *repos.App,
) {

	r := router{
		app: app,
	}

	app.Srv.HandleFunc(POST+"/login", r.handleLoginExchange)
	app.Srv.Handle(GET+"/ping", r.app.Auth.Authorize(r.handlePing))
}
