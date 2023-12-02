package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)
func (app *config) routes() http.Handler {

	mux := chi.NewRouter()
	mux.Post("sendotp", app.sendOTP)
	return mux
}
