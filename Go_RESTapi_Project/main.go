package main

//All imports

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Data Part
type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var Products []Product //Products array, containing all products

/*--- Data Part Ends ---*/

//API Part

// handler 1
func displayAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: display all products.")
	json.NewEncoder(w).Encode(Products) //Encode products and send them as a JSON response.
}

// homepage handler
func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: homepage")
	fmt.Fprintf(w, "Welcome to the Products homepage!")
}

// handler 2
func displayProduct(w http.ResponseWriter, r *http.Request) {

	log.Println("Endpoint hit: displayProduct. Product to be displayed as per URL:- ", r.URL.Path)
	key, err := strconv.Atoi(r.URL.Path[len("product/")+1:]) //get the product id from URL
	if err != nil {
		fmt.Fprintf(w, "The product with id %s doesn't exist", r.URL.Path[len("product/")+1:])
	} else {
		found := false
		for _, product := range Products {
			if product.Id == key {
				found = true
				json.NewEncoder(w).Encode(product) //when match is found write the product ot the screen
				break
			}
		}
		if !found {
			fmt.Fprintf(w, "The product with id %s doesn't exist", r.URL.Path[len("product/")+1:])
		}
	}

}

// All routes listed with there handlers
// Also ListenAndServe declared to spin up web-server
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	http.HandleFunc("/products", displayAll)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/product/", displayProduct)
	http.ListenAndServe(":8000", nil)
}

/*--- API Part ends ---*/

// Application entrypoint
func main() {

	Products = []Product{
		{Id: 1, Name: "Apple", Quantity: 100, Price: 10.0},
		{Id: 2, Name: "Mango", Quantity: 250, Price: 45.0},
		{Id: 3, Name: "Banana", Quantity: 200, Price: 12.00},
		{Id: 4, Name: "Pineapple", Quantity: 55, Price: 25.00},
	}

	handleRequests()

}
