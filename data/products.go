package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	InternalId  string  `json:"internalId"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

func (p *Product) FromJson(ir io.Reader) error {
	e := json.NewDecoder(ir)
	return e.Decode(p)
}

var productList = []*Product{
	{
		Id:          1,
		Name:        "Laptop",
		Description: "High-performance laptop with the latest specifications.",
		Price:       999.99,
		InternalId:  "laptop-001",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	},
	{
		Id:          2,
		Name:        "Smartphone",
		Description: "Feature-rich smartphone with a stunning display.",
		Price:       599.99,
		InternalId:  "smartphone-001",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	},
	{
		Id:          3,
		Name:        "Tablet",
		Description: "Portable tablet for entertainment and productivity on-the-go.",
		Price:       299.99,
		InternalId:  "tablet-001",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	},
	{
		Id:          4,
		Name:        "Headphones",
		Description: "High-quality headphones for immersive audio experience.",
		Price:       149.99,
		InternalId:  "headphones-001",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	},
	{
		Id:          5,
		Name:        "Smartwatch",
		Description: "Sleek smartwatch with fitness tracking and notifications.",
		Price:       199.99,
		InternalId:  "smartwatch-001",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	},
	{
		Id:          6,
		Name:        "Camera",
		Description: "Professional-grade camera for capturing stunning photos and videos.",
		Price:       799.99,
		InternalId:  "camera-001",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	},
	{
		Id:          7,
		Name:        "Printer",
		Description: "High-speed printer for fast and reliable printing.",
		Price:       249.99,
		InternalId:  "printer-001",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	},
	{
		Id:          8,
		Name:        "Wireless Speaker",
		Description: "Portable wireless speaker for enjoying music anywhere.",
		Price:       79.99,
		InternalId:  "speaker-001",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	},
	{
		Id:          9,
		Name:        "External Hard Drive",
		Description: "High-capacity external hard drive for secure storage.",
		Price:       129.99,
		InternalId:  "hdd-001",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	},
	{
		Id:          10,
		Name:        "Gaming Console",
		Description: "Next-generation gaming console for immersive gaming experiences.",
		Price:       399.99,
		InternalId:  "console-001",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	},
}

func GetProducts() []*Product {
	return productList
}

func AddProduct(product *Product) *Product {
	product.Id = len(productList) + 1
	productList = append(productList, product)
	return productList[len(productList)-1]
}

func UpdateProduct(id int) *Product {
	return productList[id]
}
