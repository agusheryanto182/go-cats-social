package main

import (
	"log"
	"net/http"
	"time"

	"github.com/agusheryanto182/go-social-media/internal/app"
	"github.com/agusheryanto182/go-social-media/internal/config"
	"github.com/agusheryanto182/go-social-media/internal/controller"
	"github.com/agusheryanto182/go-social-media/internal/repository"
	"github.com/agusheryanto182/go-social-media/internal/service"
	"github.com/agusheryanto182/go-social-media/utils/hash"
	"github.com/agusheryanto182/go-social-media/utils/jwt"
	"github.com/go-playground/validator/v10"
)

func main() {
	cfg := config.NewConfig()
	db := config.InitialDB(cfg)
	hash := hash.NewHash(cfg)
	jwt := jwt.NewJWT(cfg.Jwt.Secret)
	valid := validator.New()

	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo, db, hash, jwt)
	userCtrl := controller.NewUserController(userSvc, valid)

	router := app.NewRouter(userCtrl)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
