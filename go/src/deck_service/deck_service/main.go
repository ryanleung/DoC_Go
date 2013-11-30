package main

import (
	"deck_service/base"
	"flag"
	"deck_service/controllers"
	// "fmt"
	"log"
)

var port = flag.Int("port", 8000, "Theport on which to listen for requests")
var databaseSSL = flag.Bool(
	"db_ssl",
	false,
	"Whether to attempt to use SSL")
var databaseUri = flag.String(
	"db_uri",
	"postgres",
	"URI of postgres database to connect with")

func main() {
	flag.Parse()

	log.Println("Creating Game Manager...")
	gm := base.NewGameManager()
	log.Println("Game Manager Initialized.")

	log.Println("Starting up json service")
	jsonService := controllers.JsonService{gm}
	jsonService.Serve(*port)
}
