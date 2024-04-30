package controller

import "net/http"

type MatchController interface {
	Match(w http.ResponseWriter, r *http.Request)
	GetMatch(w http.ResponseWriter, r *http.Request)
	Approve(w http.ResponseWriter, r *http.Request)
	Reject(w http.ResponseWriter, r *http.Request)
}
