package authRouter

import (
	"net/http"
)

func (r router) handlePing(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
