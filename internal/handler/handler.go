package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/captain-corgi/go-graceful-shutdown-multiple-services/pkg/restful"
)

func testGracefulShutDown(res http.ResponseWriter, req *http.Request) {
	time.Sleep(10 * time.Second)
	log.Println("testGracefulShutdown job completed")
	restful.ResponseWithJSON(res, 200, map[string]interface{}{"status": "completed"})
}

func New(r *mux.Router) {
	r.HandleFunc("/test-graceful-shutdown", testGracefulShutDown).Methods(http.MethodGet)
}
