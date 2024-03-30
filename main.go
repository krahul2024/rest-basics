package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-basics/handlers"
	"time"
)

func main() {
	l := log.New(os.Stdout, "\nExample API\n", log.LstdFlags) // logger initialization
	indexHandler := handlers.NewIndex(l)
	productsHandler := handlers.NewProducts(l)

	mux := http.NewServeMux()
	mux.Handle("/", indexHandler)
	mux.Handle("/products/", productsHandler)

	server := &http.Server{
		Addr:         ":3300",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	log.Printf("Starting the server at the PORT : %v", 3300)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal("Error starting the server!\n", err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)
	sig := <-sigChannel
	l.Printf("Termination request received(%v), shutting down gracefully!", sig)

	// shutdown enabling with absolute time
	/* duration := time.Now().Add((30 * time.Second))
	timeoutContext, cancel := context.WithDeadline(context.Background(), duration)
	defer cancel()                  // this function is called after the main function reaches the return statment or the end of its scope to cancel the context else it would continue to consume the resources
	server.Shutdown(timeoutContext) // shuts down the server after specified time no matter whether the operations are still going on
	*/

	// graceful shutdown : with timeout, in this case server would shut down if all the requests are completed within the specified time, if there are ongoing requests it would wait for them to finish.
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}

// read more about graceful shutdown implementation
// read more about the json serialization
