package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-basics/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(res, req)
	}
	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(res http.ResponseWriter, req *http.Request) {
	productsList := data.GetProducts()
	jsonData, err := json.Marshal(productsList) // we can also use encoder to encode which is faster

	if err != nil {
		http.Error(res, "An error occured getting the list of products!", http.StatusInternalServerError)
	}
	res.Write(jsonData)
}
