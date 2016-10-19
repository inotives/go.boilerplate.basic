package routers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// TestRoute ...
func TestRoute(router *mux.Router) *mux.Router {

	// first routes
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "THIS IS ROOT PATH")
	})

	return router
}
