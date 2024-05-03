package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	catRepo := repository.NewCatRepository(db)
	matchRepo := repository.NewMatchRepository(db)

	userSvc := service.NewUserService(userRepo, db, hash, jwt)
	catSvc := service.NewCatService(catRepo, db)
	matchSvc := service.NewMatchService(db, matchRepo, catRepo)

	userCtrl := controller.NewUserController(userSvc, valid)
	catCtrl := controller.NewCatController(catSvc, valid, matchSvc)
	matchCtrl := controller.NewMatchController(matchSvc, catSvc, valid)

	router := app.NewRouter(userCtrl, catCtrl, userSvc, jwt, matchCtrl)

	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	go func() { // create a goroutine for servers
		log.Fatal(srv.ListenAndServe())
	}()

	quit := make(chan os.Signal, 1)   // create channel that can receive signal
	signal.Notify(quit, os.Interrupt) // set quit that can receive signal from os.Interrupt , when program receive signal os.Interrupt,
	// it will send to quit, and program will exit
	<-quit // block program until receive signal
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // create context with timeout and cancel
	defer cancel()                                                          // should be running even if program exit, error or panic
	if err := srv.Shutdown(ctx); err != nil {                               // shutdown server with context
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

}
