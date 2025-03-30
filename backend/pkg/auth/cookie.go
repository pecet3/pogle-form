package auth

import (
	"net/http"
	"time"
)

func (a *Auth) SetCookies(w http.ResponseWriter, c string) {
	cookie := http.Cookie{
		Name:     "auth",
		Value:    c,
		Expires:  time.Now().Add(time.Hour * 192),
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}

	http.SetCookie(w, &cookie)
}
