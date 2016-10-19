package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"template.basic/routers"

	"github.com/urfave/negroni"
)

func main() {

	port := 1920 // might need to change this to config / settings ...

	router := routers.InitRoutes()

	n := negroni.New()
	n.Use(negroni.NewLogger()) // logger onto console

	n.UseHandler(router)

	srv := &http.Server{
		Handler: n,
		Addr:    "127.0.0.1:" + strconv.Itoa(port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Server serving at port:: %v \n", port)
	log.Fatal(srv.ListenAndServe())
}
