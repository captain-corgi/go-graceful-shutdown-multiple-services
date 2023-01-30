package router

import (
	"github.com/gorilla/mux"

	"github.com/captain-corgi/go-graceful-shutdown-multiple-services/internal/handler"
)

func New() *mux.Router {
	router := mux.NewRouter()
	apiV1Router := router.PathPrefix("/api/v1").Subrouter()
	handler.New(apiV1Router)
	return router
}
