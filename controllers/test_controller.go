package controllers

import "net/http"

// HelloController ...
func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Heeeellllo!!"))
}
