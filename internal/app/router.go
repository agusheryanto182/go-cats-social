package app

import (
	"github.com/agusheryanto182/go-social-media/internal/controller"
	"github.com/gorilla/mux"
)

func NewRouter(userCtrl controller.UserController) *mux.Router {
	r := mux.NewRouter()

	user := r.PathPrefix("/v1/user").Subrouter()
	user.HandleFunc("/register", userCtrl.Register).Methods("POST")
	user.HandleFunc("/login", userCtrl.Login).Methods("POST")

	return r
}
