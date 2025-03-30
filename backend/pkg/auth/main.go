package auth

import (
	"pogle-form/backend/pkg/auth/login"
)

type Auth struct {
	LoginSession *login.Login
}

func New() *Auth {
	a := &Auth{
		LoginSession: login.New(),
	}

	return a
}
