package controller

import "net/http"

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request)
}
