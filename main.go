package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("There was an error running the server")
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the index page!"))

}
