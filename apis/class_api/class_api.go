package class_api

import (
	"encoding/json"
	"goworkspace/config"
	"goworkspace/entities"
	"goworkspace/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		classModel := models.ClassModel{
			Db: db,
		}
		classes, err2 := classModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJSON(response, http.StatusOK, classes)
		}
	}
}
func Addingstudent(response http.ResponseWriter, r *http.Request) {

	var information entities.Class
	json.NewDecoder(r.Body).Decode(&information)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		classModel := models.ClassModel{
			Db: db,
		}
		//respondWithJSON(response, http.StatusOK, information)
		err1 := classModel.Addingstudent(&information)
		if err1 != nil {
			respondWithError(response, http.StatusBadRequest, err1.Error())
		} else {
			respondWithJSON(response, http.StatusOK, information)
		}

	}
}
func Updateemployee(response http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]
	i, _ := strconv.Atoi(key)
	var information entities.Class
	json.NewDecoder(r.Body).Decode(&information)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		classModel := models.ClassModel{
			Db: db,
		}
		err1 := classModel.Updateemployee(i, &information)
		if err1 != nil {
			respondWithError(response, http.StatusBadRequest, err1.Error())
		} else {
			respondWithJSON(response, http.StatusOK, "updated")
		}
	}
}

func Updateparticularemployee(response http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]
	i, _ := strconv.Atoi(key)
	var information entities.Class
	json.NewDecoder(r.Body).Decode(&information)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		classModel := models.ClassModel{
			Db: db,
		}
		err1 := classModel.Updateparticularemployee(i, &information)
		if err1 != nil {
			respondWithError(response, http.StatusBadRequest, err1.Error())
		} else {
			respondWithJSON(response, http.StatusOK, "updated")
		}
	}
}

func Deletemployee(response http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["id"]
	i, _ := strconv.Atoi(key)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		classModel := models.ClassModel{
			Db: db,
		}
		err1 := classModel.Deletemployee(i)
		if err1 != nil {
			respondWithError(response, http.StatusBadRequest, err1.Error())
		} else {
			respondWithJSON(response, http.StatusOK, "deleted")
		}
	}
}
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
