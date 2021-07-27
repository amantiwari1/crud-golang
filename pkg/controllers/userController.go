package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bitly/go-simplejson"

	"github.com/gorilla/mux"

	"strconv"

	"github.com/amantiwari1/crud-golang/pkg/models"
	"github.com/amantiwari1/crud-golang/pkg/utils"
)

var NewBook models.User

func GetUser(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllUser()
	res, _ := json.Marshal(newBooks)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetUserById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.User{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateUser()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	userExist, _ := models.GetUserById(ID)

	if userExist.ID == 0 {
		json := simplejson.New()
		json.Set("message", "user is not exist")

		payload, _ := json.MarshalJSON()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(payload)

		return
	}

	if err != nil {
		fmt.Println("error while parsing")
	}
	models.DeleteUser(ID)

	json := simplejson.New()
	json.Set("message", "user deleted")

	payload, _ := json.MarshalJSON()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var updateBook = &models.User{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	userExist, _ := models.GetUserById(ID)

	if userExist.ID == 0 {
		json := simplejson.New()
		json.Set("message", "user is not exist")

		payload, _ := json.MarshalJSON()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(payload)

		return
	}
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetUserById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Email != "" {
		bookDetails.Email = updateBook.Email
	}
	if updateBook.Gender != "" {
		bookDetails.Gender = updateBook.Gender
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
