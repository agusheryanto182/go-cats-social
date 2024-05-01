package app

import (
	"github.com/agusheryanto182/go-social-media/internal/controller"
	"github.com/agusheryanto182/go-social-media/internal/middleware"
	"github.com/agusheryanto182/go-social-media/internal/service"
	"github.com/agusheryanto182/go-social-media/utils/jwt"
	"github.com/gorilla/mux"
)

func NewRouter(
	userCtrl controller.UserController,
	catCtrl controller.CatController,
	userSvc service.UserService,
	jwtSvc jwt.IJwt,
	matchCtrl controller.MatchController,
) *mux.Router {

	r := mux.NewRouter()

	// user
	user := r.PathPrefix("/v1/user").Subrouter()
	user.HandleFunc("/register", userCtrl.Register).Methods("POST")
	user.HandleFunc("/login", userCtrl.Login).Methods("POST")

	// cat
	cat := r.PathPrefix("/v1/cat").Subrouter()
	cat.Use(middleware.NewAuthMiddleware(userSvc, jwtSvc).Protected)
	cat.HandleFunc("", catCtrl.Create).Methods("POST")
	cat.HandleFunc("", catCtrl.GetCat).Methods("GET")
	cat.HandleFunc("/{id}", catCtrl.Update).Methods("PUT")
	cat.HandleFunc("/{id}", catCtrl.Delete).Methods("DELETE")

	cat.HandleFunc("/match", matchCtrl.Match).Methods("POST")
	cat.HandleFunc("/match", matchCtrl.GetMatch).Methods("GET")
	cat.HandleFunc("/match/approve", matchCtrl.Approve).Methods("POST")
	cat.HandleFunc("/match/reject", matchCtrl.Reject).Methods("POST")
	cat.HandleFunc("/match/{id}", matchCtrl.DeleteTheMatch).Methods("DELETE")

	return r
}
