package app

import (
	"github.com/agusheryanto182/go-social-media/internal/controller"
	"github.com/agusheryanto182/go-social-media/internal/middleware"
	"github.com/agusheryanto182/go-social-media/internal/service"
	"github.com/agusheryanto182/go-social-media/utils/jwt"
	"github.com/gorilla/mux"
)

func NewRouter(userCtrl controller.UserController, catCtrl controller.CatController, userSvc service.UserService, jwtSvc jwt.IJwt) *mux.Router {
	r := mux.NewRouter()

	user := r.PathPrefix("/v1/user").Subrouter()
	user.HandleFunc("/register", userCtrl.Register).Methods("POST")
	user.HandleFunc("/login", userCtrl.Login).Methods("POST")

	cat := r.PathPrefix("/v1/cat").Subrouter()
	cat.Use(middleware.NewAuthMiddleware(userSvc, jwtSvc).Protected)
	cat.HandleFunc("", catCtrl.Create).Methods("POST")
	cat.HandleFunc("", catCtrl.GetCat).Methods("GET")
	cat.HandleFunc("/{id}", catCtrl.Update).Methods("PUT")

	return r
}
