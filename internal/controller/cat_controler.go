package controller

import "net/http"

type CatController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetCat(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
