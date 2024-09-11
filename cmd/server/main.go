package main

import (
	"creds-update/config"
	"creds-update/handlers"
	"creds-update/resources"
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB
var router http.Handler
var routeHandler *handlers.RouteHandler

func main() {
	db = config.GetDatabase()
	routeHandler = handlers.NewRouteHandler(db)
	router = config.GetRoutes(routeHandler)
	err := resources.LoadViews()
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":3000", router)
}
