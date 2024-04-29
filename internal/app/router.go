package app

import (
	"github.com/agusheryanto182/go-social-media/internal/controller"
	"github.com/gorilla/mux"
)

func NewRouter(userCtrl controller.UserController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/v1/user/register", userCtrl.Register).Methods("POST")

	return r
}
