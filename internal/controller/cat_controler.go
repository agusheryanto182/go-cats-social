package controller

import "net/http"

type CatController interface {
	Create(w http.ResponseWriter, r *http.Request)
}
