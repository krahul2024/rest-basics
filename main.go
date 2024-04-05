package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"rss-agg/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           30000,
	}))

	statusRouter := chi.NewRouter()
	statusRouter.Get("/status", utils.ReportStatus)
	router.Mount("/", statusRouter)

	// index routes
	router.Get("/", IndexGetHandler)
	router.Post("/", IndexPostHandler)

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%v", PORT),
	}

	log.Printf("Starting the server at PORT : %v\n", PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("There was an error starting the server at the PORT : ", PORT, "\n", err)
	}

}

func IndexGetHandler(res http.ResponseWriter, req *http.Request) {
	response := struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{"This is some message", 200}
	jsonRes, err := json.Marshal(response)
	if err != nil {
		http.Error(res, "Something went wrong!", http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonRes)
}

func IndexPostHandler(res http.ResponseWriter, req *http.Request) {
	var reqBody map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		log.Printf("There was an error!, %v", err)
		res.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Printf("%+v\n", reqBody)
	fmt.Println(reqBody["age"])

	utils.RespondWithJson(res, http.StatusOK, reqBody)
}
