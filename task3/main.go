package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Name  string
	ID    string
	Qty   int
	Price float64
}

var Products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit by muxRouter: Homepage")
	fmt.Println("Welcome to Homepage served by gorilla/mux module")
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit by muxRouter: getProducts")
	vars := mux.Vars(r)
	key := vars["ID"]
	for _, product := range Products {
		if string(product.ID) == key {
			json.NewEncoder(w).Encode(product)
		}
	}
}

func returnProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnProducts")
	json.NewEncoder(w).Encode(Products)
}

func handleRequest() {
	//creates a new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/homepage", homepage)
	myRouter.HandleFunc("/products", returnProducts)
	myRouter.HandleFunc("/product/{ID}", getProducts)
	http.ListenAndServe("localhost:10030", myRouter)
}

func main() {
	Products = []Product{
		Product{Name: "pencil", ID: "1", Qty: 2, Price: 1.5},
		Product{Name: "sharpener", ID: "2", Qty: 1, Price: 2.5},
	}
	handleRequest()
}
