package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/tulza/go-api/internal/middleware"
)


func Handler(router *chi.Mux) {
	// middleware
	router.Use(chimiddle.StripSlashes) // strip slash e.g. account/coin/ to account/coin

	router.Route("/account", func(router chi.Router){
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}