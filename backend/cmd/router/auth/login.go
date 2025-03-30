package authRouter

import (
	"net/http"
	"os"
	"pogle-form/backend/data/dtos"

	"github.com/pecet3/logger"
)

func (r router) handleLoginExchange(w http.ResponseWriter, req *http.Request) {
	dto := &dtos.Login{}
	if err := r.app.Dto.Get(req, dto); err != nil {
		logger.Error(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	name := os.Getenv("ADMIN_NAME")
	password := os.Getenv("ADMIN_PASSWORD")
	if dto.Name != name || dto.Password != password {
		logger.Error("Wrong credentials")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	session := r.app.Auth.LoginSession.Add()
	r.app.Auth.SetCookies(w, session.Cookie)

}
