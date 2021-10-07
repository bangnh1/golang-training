package routes

import (
	"encoding/json"
	"github.com/bangnh1/golang-training/07/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

func ReturnAllUsers(w http.ResponseWriter, r *http.Request) {

	query := r.FormValue("sort")

	users := models.ListUser()
	switch query {
	case "age":
		sort.Sort(models.ByAge(users))
	case "name":
		sort.Sort(models.ByName(users))
	default:
		sort.Sort(models.ById(users))
	}

	usersJson, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJson)
}

func ReturnSingleUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, user := range models.ListUser() {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if user.Id == idInt {
			usersJson, err := json.Marshal(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(usersJson)
		}
	}

}

func CreateNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	newUser := &models.User{}
	err := json.Unmarshal(reqBody, newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users := models.ListUser()
	users = append(users, newUser)
	usersJson, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(usersJson)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := models.ListUser()
	reqBody, _ := ioutil.ReadAll(r.Body)
	newUser := &models.User{}
	err = json.Unmarshal(reqBody, newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, user := range users {
		if user.Id == idInt {
			users[index] = newUser
		}
	}
	newUserJson, err := json.Marshal(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(newUserJson)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := models.ListUser()
	newusers := []*models.User{}
	for _, user := range users {

		if user.Id != idInt {
			newusers = append(newusers, user)
		} else {
			newUserJson, err := json.Marshal(user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(newUserJson)
		}
	}

}
