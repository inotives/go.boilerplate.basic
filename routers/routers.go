package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes ...
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router = TestRoute(router)
	return router
}
