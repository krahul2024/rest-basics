package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest-basics/data"
	"rest-basics/utils"
	"strconv"
	"strings"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// reqParams := utils.GetParams(req)
	fmt.Println(req.Method)
	if req.Method == http.MethodGet {
		p.getProducts(res, req)
		return
	}
	if req.Method == http.MethodPost {
		p.addProduct(res, req)
		return
	}
	if req.Method == http.MethodPut {
		// this is for getting the products or the url should be something like /products/id
		urlParts := strings.Split(req.URL.String(), "/")
		if len(urlParts) != 3 {
			http.Error(res, "There was an error!", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(urlParts[2])
		if err != nil {
			http.Error(res, "There was an error!", http.StatusBadRequest)
			return
		}
		p.updateProduct(id-1, res, req)
	}
}

func (p *Products) getProducts(res http.ResponseWriter, req *http.Request) {
	p.l.Println("Get Product route called!")
	productsList := data.GetProducts()
	jsonData, err := json.Marshal(productsList) // we can also use encoder to encode which is faster

	if err != nil {
		http.Error(res, "An error occured getting the list of products!", http.StatusInternalServerError)
	}
	res.Write(jsonData)
}

func (p *Products) addProduct(res http.ResponseWriter, req *http.Request) {
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

func (p *Products) updateProduct(id int, res http.ResponseWriter, req *http.Request) {
	p.l.Println("Update product route called!")
	product := data.UpdateProduct(id)
	utils.PrettyJSON(product)
	res.Write([]byte("Product updated!"))
}
