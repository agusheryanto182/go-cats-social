package controller

import "net/http"

type MatchController interface {
	Match(w http.ResponseWriter, r *http.Request)
}
