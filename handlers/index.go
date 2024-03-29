package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Index struct {
	l *log.Logger
}

func NewIndex(l *log.Logger) *Index {
	return &Index{l}
}

func (idx *Index) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	idx.l.Println("This is the index route!")
	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "An error occured!", http.StatusBadRequest)
		return
	}
	// res.Write([]byte("Welcome to the index page! " + string(data))) // use the below format since this one requires us to have the information in byte but the one below is more flexible
	fmt.Fprintf(res, "%s", []byte("Welcome to the index page! "+string(data)))
}
