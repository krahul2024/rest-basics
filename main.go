package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-basics/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	PORT := 3300
	l := log.New(os.Stdout, "Example-API\n", log.LstdFlags) // basic logger

	indexHandler := handlers.NewIndex(l)
	productsHandler := handlers.NewProducts(l)

	router := mux.NewRouter()
	getRouter := router.Methods(http.MethodGet).Subrouter()
	postRouter := router.Methods(http.MethodPost).Subrouter()
	putRouter := router.Methods(http.MethodPut).Subrouter()

	// get routes
	getRouter.HandleFunc("/", indexHandler.IndexRoute)
	getRouter.HandleFunc("/products", productsHandler.GetProducts)

	// post routes
	postRouter.HandleFunc("/products", productsHandler.AddProduct)

	// put routes
	putRouter.HandleFunc("/products/{id:[0-9]+}", productsHandler.UpdateProduct)

	server := &http.Server{
		Addr:     fmt.Sprintf(":%v", PORT),
		Handler:  router,
		ErrorLog: l,
	}
	log.Printf("Starting the server at PORT : %v", PORT)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal("Error starting the server!\n", err)
		}
	}()

	// send the notification before shutdown
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)
	sig := <-sigChannel
	l.Printf("Termination request received(%v), shutting down gracefully!", sig)

	// graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
