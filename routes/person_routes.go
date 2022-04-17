package routes

import (
	"github.com/gorilla/mux"
	"github.com/jeremiascardoso00/demo-crud-api-rest-go/controllers"
)

func SetPersonRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("person/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete", controllers.Delete).Methods("DELETE")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")

}
