package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Name  string  `json: "Name"`
	ID    string  `json: "ID"`
	Qty   int     `json: "Qty"`
	Price float64 `json: "Price"`
}

var Products []Product

// homepage handler
func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit by muxRouter: Homepage")
	w.Header().Set("Content-type", "text/html")
	w.Write([]byte("<h1 style='color: #005500'>Welcome to Homepage !</h2>\n"))
}

// get products  with specific IDs
func getProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit by muxRouter: getProducts")
	vars := mux.Vars(r)
	key := vars["ID"]
	for _, product := range Products {
		if string(product.ID) == key {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(product)
		}
	}
}

// returns all available Products
func returnProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	log.Println("Endpoint hit: returnProducts")
	json.NewEncoder(w).Encode(Products)
}

func handleRequest() {
	//creates a new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	//match route URLs to handlers
	myRouter.HandleFunc("/homepage", homepage)
	myRouter.HandleFunc("/products", returnProducts)
	myRouter.HandleFunc("/product/{ID}", getProducts)

	http.ListenAndServe("localhost:10050", myRouter)
}

func main() {
	Products = []Product{
		Product{Name: "pencil", ID: "1", Qty: 2, Price: 1.5},
		Product{Name: "sharpener", ID: "2", Qty: 1, Price: 2.5},
		Product{Name: "ruler", ID: "3", Qty: 5, Price: 2},
	}
	handleRequest()
}
