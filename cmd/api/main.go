package main

import (
	"fmt"
	"net/http"

	"github.com/tulza/go-api/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var router *chi.Mux = chi.NewRouter()
	handlers.Handler(router)

	fmt.Println("Starting GO API service...")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Error(err)
	}
}
