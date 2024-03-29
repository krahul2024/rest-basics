package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello!")

	http.HandleFunc("/", indexFunction)

	log.Fatal(http.ListenAndServe(":3300", nil))
}

func indexFunction(res http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "An error occured", http.StatusBadRequest)
		return
	}
	log.Printf("%s", data)
	res.Write([]byte("This is index page..." + string(data)))
}
