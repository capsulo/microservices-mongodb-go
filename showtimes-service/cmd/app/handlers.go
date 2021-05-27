package app

import (
	"cinema.cassia.io/showtimes/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func (app *Application) all(w http.ResponseWriter, r *http.Request) {
	showtimes, err := app.Showtimes.All()
	if err != nil {
		app.serverError(w, err)
	}

	b, err := json.Marshal(showtimes)
	if err != nil {
		app.serverError(w, err)
	}
	app.InfoLog.Println("Showtimes have been listed")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *Application) findById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	m, err := app.Showtimes.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.InfoLog.Println("Showtime not found")
			return
		}
		app.serverError(w, err)
	}

	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}
	app.InfoLog.Println("Have been found a showtime")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *Application) findByDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]

	m, err := app.Showtimes.FindByDate(date)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.InfoLog.Println("Showtime not found")
			return
		}
		app.serverError(w, err)
	}
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
	}
	app.InfoLog.Println("Have been found a showtime")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *Application) insert(w http.ResponseWriter, r *http.Request) {
	var m models.ShowTime
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
	}
	m.CreatedAt = time.Now().String()
	insertResult, err := app.Showtimes.Insert(m)
	if err != nil {
		app.serverError(w, err)
	}
	app.InfoLog.Printf("New showtime have been created id=%s", insertResult.InsertedID)
}

func (app *Application) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	deleteResult, err := app.Showtimes.Delete(id)
	if err != nil {
		app.serverError(w, err)
	}
	app.InfoLog.Printf("Have been eliminated %d showtime(s)", deleteResult.DeletedCount)
}

