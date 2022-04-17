package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jeremiascardoso00/demo-crud-api-rest-go/commons"
	"github.com/jeremiascardoso00/demo-crud-api-rest-go/models"
	"log"
	"net/http"
)

func GetAll(writer http.ResponseWriter, request *http.Request) {
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

func Get(writer http.ResponseWriter, request *http.Request) {
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

func Save(writer http.ResponseWriter, request *http.Request) {
	person := models.Person{}
	db, psgDB := commons.GetConnection()
	defer func() {
		psgDB.Close()
	}()

	err := json.NewDecoder(request.Body).Decode(person)
	if err != nil {
		log.Fatal("error in json decoder", err)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	err = db.Create(&person).Error
	if err != nil {
		log.Fatal("error in db create", err)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(person)
	if err != nil {
		log.Fatal("error in json marshal", err)
		return
	}
	commons.SendResponse(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	person := models.Person{}
	db, psgDB := commons.GetConnection()
	defer func() {
		psgDB.Close()
	}()
	id := mux.Vars(request)["id"]
	db.Find(&person, id)
	if person.ID > 0 {
		err := db.Delete(&person).Error
		if err != nil {
			log.Fatal("error in db delete", err)
			commons.SendError(writer, http.StatusInternalServerError)
			return
		}
		commons.SendResponse(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
