package main

import (
	"cinema.cassia.io/showtimes/cmd/app"
	"cinema.cassia.io/showtimes/pkg/models/mongodb"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	serverAddr := flag.String("serverAddr", "localhost", "HTTP Network server address")
	serverPort := flag.Int("serverPort", 4002, "Http Server network port")
	mongoURI := flag.String("mongoURI", "mongodb://localhost:27017", "Database hostname url")
	mongoDB := flag.String("mongoDB", "showtimes", "Database name")
	enableCredentials := flag.Bool("enableCredentials", false, "Enable the use of credentials for mongo connection")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	connection := options.Client().ApplyURI(*mongoURI)
	if *enableCredentials {
		connection.Auth = &options.Credential{
			Username: os.Getenv("MONGODB_USERNAME"),
			Password: os.Getenv("MONGODB_PASSWORD"),
		}
	}

	client, err := mongo.NewClient(connection)
	if err != nil {
		errLog.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		errLog.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	infoLog.Printf("Database connection established")

	// Initialize a new instance of application containing the dependencies.
	a := &app.Application{
		InfoLog:  infoLog,
		ErrorLog: errLog,
		Showtimes: &mongodb.ShowTimeModel{
			C: client.Database(*mongoDB).Collection("showtimes"),
		},
	}

	// Initialize a new http.Server struct.
	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	srv := &http.Server{
		Addr:         serverURI,
		ErrorLog:     errLog,
		Handler:      a.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", serverURI)
	err = srv.ListenAndServe()
	errLog.Fatal(err)
}
