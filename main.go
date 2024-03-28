package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not set in the environment variables!")
	}

	router := mux.NewRouter()

	// index route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Request Method:", r.Method)
		fmt.Println("Request URL:", r.URL.String())
		fmt.Println("Request Headers:")
		for name, values := range r.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", name, value)
			}
		}
		fmt.Println("Request body\n", r.Body)

		message := map[string]interface{}{
			"message": "Welcome to the index page",
			"status":  200,
			"success": true,
		}
		jsonResponse, err := json.Marshal(message)
		if err != nil {
			http.Error(w, "An error occured parsing JSON!", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server running on the port : %v", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}
