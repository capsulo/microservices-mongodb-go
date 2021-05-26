package main

import (
	"cinema.cassia.io/users/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	users, err := app.users.All()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(users)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Println("Users have been listed")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.users.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("User not found")
			app.clientError(w, http.StatusBadRequest)
			return
		}
		app.serverError(w, err)
	}
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Println("Have been found a user")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		app.serverError(w, err)
	}
	insertResult, err := app.users.Insert(u)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("New user have been created , id=%s", insertResult.InsertedID)
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteResult, err := app.users.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}
	app.infoLog.Printf("Have been eliminated %d user(s)", deleteResult.DeletedCount)
}
