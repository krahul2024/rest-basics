package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// mux.HandleFunc("/", indexHandler)
	// mux.HandleFunc("/about", aboutHandler)
	// mux.HandleFunc("/post", postHandler)

	mux.HandleFunc("/", handleRequest(indexHandler, http.MethodGet, http.MethodPost))
	mux.HandleFunc("/about", handleRequest(aboutHandler, http.MethodGet))
	mux.HandleFunc("/post", handleRequest(postHandler, http.MethodPost))

	server := &http.Server{
		Addr:    ":3300",
		Handler: mux,
	}

	log.Printf("Starting the server at PORT : 3300\n")

	server.ListenAndServe()
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "An error occured!\n", http.StatusBadRequest)
		return
	}
	res.Write([]byte("This is the index route!\n" + string(data)))
}

func aboutHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("This is the about route!\n"))
}

func postHandler(res http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "An error occured!\n", http.StatusBadRequest)
		return
	}
	res.Write([]byte("This is the post route!\n" + string(data)))
}

/* This function removes the repetitive task of checking method type in each of the xhandler
functions, otherwise we need to check what is the request method type and then perform
the action accordingly.
*/

func handleRequest(handler http.HandlerFunc, allowedMethods ...string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		for _, method := range allowedMethods {
			if req.Method == method {
				handler(res, req)
				return
			}
		}
	}
}

// Task to complete
/*
1. Perform actual get and post methods using this approach
2. Perform database related operations using this approach
3. Perform authentication and cookie setting using this approach
*/
