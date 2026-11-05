package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

var product = []Product{
	{1, "phone", 4000, "electronic"},
	{2, "AC", 39054, "electronic"},
	{3, "TV", 56789, "electronic"},
	{5, "Washing machine", 46899, "electronic"},
}

// GET :  Fetch all products
func getProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowd ", 400)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// GET :Fetch product by id
func get_by_id(w http.ResponseWriter, r *http.Request) {

	//by qurey
	id := r.URL.Query().Get("id")
	for _, p := range product {
		if strconv.Itoa(p.Id) == id {
			w.Header().Set("Content-type", "applictio/json")
			json.NewEncoder(w).Encode(p)

		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

//POST : add a product

func addProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allwod", 405)
		return
	}
	var NewProduct Product
	err := json.NewDecoder(r.Body).Decode(&NewProduct)
	if err != nil {
		http.Error(w, "invalid ", 400)
		return
	}
	product = append(product, NewProduct)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

// Main Function

func main() {
	http.HandleFunc("/all-product", getProduct)
	http.HandleFunc("/product", get_by_id)
	http.HandleFunc("/add-product", addProduct)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
