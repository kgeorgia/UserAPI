package main

import (
	"log"
	"net/http"
	"os"
	"refactoring/internal/app"
	"refactoring/internal/controller/handler"
)

const store = `users.json`

func main() {

	port := ":3333"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	r := app.NewRouter()

	h, err := handler.New(store)
	if err != nil {
		log.Fatal("handler: ", err.Error())
	}

	app.Route(r, h)

	err = http.ListenAndServe(port, r)
	log.Fatal(err)
}
