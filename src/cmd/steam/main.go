package main

import (
	"api"
	"fmt"
	"net/http"
	"os"

	"steam"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	steam.WebApiKey = os.Getenv("STEAM_WEB_API_KEY")

	if port == "" {
		port = "8080"
	}

	if steam.WebApiKey == "" {
		fmt.Fprintf(os.Stderr, "Steam Web Api Key not defined")
		os.Exit(1)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", api.StreamsHandler)

	var handler http.Handler
	handler = handlers.CORS()(r)
	handler = handlers.CompressHandler(handler)

	fmt.Fprintf(os.Stdout, "Server launching on port %s\n", port)

	http.ListenAndServe(":"+port, handler)
}
