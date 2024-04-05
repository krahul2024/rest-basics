package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-basics/data"
	"rest-basics/utils"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Get Product route called!")
	productsList := data.GetProducts()
	jsonData, err := json.Marshal(productsList) // we can also use encoder to encode which is faster

	if err != nil {
		http.Error(res, "An error occured getting the list of products!", http.StatusInternalServerError)
	}
	res.Write(jsonData)
}

func (p *Products) AddProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Add product route called!")
	product := &data.Product{}
	err := product.FromJson(req.Body)
	if err != nil {
		http.Error(res, "There was an error!", http.StatusBadRequest)
		return
	}
	addedProduct := data.AddProduct(product)
	utils.PrettyJSON(addedProduct)
	jsonData, err := json.Marshal(addedProduct)
	if err != nil {
		http.Error(res, "There was an error!", http.StatusBadRequest)
		return
	}
	res.Write(jsonData)
}

func (p *Products) UpdateProduct(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	id -= 1
	p.l.Println("Update product route called!")
	product := data.UpdateProduct(id)
	utils.PrettyJSON(product)
	res.Write([]byte("Product updated!"))
}
