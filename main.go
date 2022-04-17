package main

import (
	"github.com/gorilla/mux"
	"github.com/jeremiascardoso00/demo-crud-api-rest-go/commons"
	"github.com/jeremiascardoso00/demo-crud-api-rest-go/routes"
	"log"
	"net/http"
)

func main() {
	var port = "9000"
	commons.Migrate()
	router := mux.NewRouter()
	routes.SetPersonRoutes(router)
	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Println("Server up in port: " + port)
	log.Println(server.ListenAndServe())
}
