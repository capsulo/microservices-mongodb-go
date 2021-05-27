package app

import (
	"cinema.cassia.io/showtimes/pkg/models/mongodb"
	"log"
)

type Application struct {
	ErrorLog  *log.Logger
	InfoLog   *log.Logger
	Showtimes *mongodb.ShowTimeModel
}
