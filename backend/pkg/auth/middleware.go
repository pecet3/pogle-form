package auth

import (
	"net/http"
	"time"

	"github.com/pecet3/logger"
)

func (as *Auth) getToken(r *http.Request) (string, error) {
	cookie, err := r.Cookie("auth")
	if err != nil || cookie.Value == "" {
		return "", err
	}

	return cookie.Value, nil
}

func (as *Auth) Authorize(next http.HandlerFunc, roles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := as.getToken(r)
		if err != nil {
			logger.Error(err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		ls, exists := as.LoginSession.Get(cookie)
		if !exists {
			logger.Error("login sessions doesn't exist")
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		if ls.Expiry.After(time.Now()) {
			logger.Error("expired session")
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
