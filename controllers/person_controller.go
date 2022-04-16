package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jeremiascardoso00/demo-crud-api-rest-go/commons"
	"github.com/jeremiascardoso00/demo-crud-api-rest-go/models"
	"log"
	"net/http"
)

func getAll(writer http.ResponseWriter, request http.Request) {
	persons := []models.Person{}
	db, psgDB := commons.GetConnection()

	defer func() {
		psgDB.Close()
	}()
	db.Find(&persons)
	json, err := json.Marshal(persons)
	if err != nil {
		log.Fatal("error in json marshal", err)
		return
	}
	commons.SendResponse(writer, http.StatusOK, json)
}

func get(writer http.ResponseWriter, request *http.Request) {
	person := models.Person{}
	id := mux.Vars(request)["id"]
	db, psgDB := commons.GetConnection()
	defer func() {
		psgDB.Close()
	}()

	db.Find(&person, id)
	if person.ID > 0 {
		json, err := json.Marshal(person)
		if err != nil {
			log.Fatal("error in json marshal", err)
			return
		}
		commons.SendResponse(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
