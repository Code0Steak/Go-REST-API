package main

//All imports

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Data Part
type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var Products []Product //Products array, containing all products

/*--- Data Part Ends ---*/

//API Part

//handler 1
func displayAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: display all products.")
	json.NewEncoder(w).Encode(Products) //Encode products and send them as a JSON response.
}

//homepage handler
func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: homepage")
	fmt.Fprintf(w, "Welcome to the Products homepage!")
}

func handleRequests() {
	http.HandleFunc("/products", displayAll)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8000", nil)
}

/*--- API Part ends ---*/

//Application entrypoint
func main() {

	Products = []Product{
		{Id: 1, Name: "Apple", Quantity: 100, Price: 10.0},
		{Id: 2, Name: "Mango", Quantity: 250, Price: 45.0},
		{Id: 3, Name: "Banana", Quantity: 200, Price: 12.00},
		{Id: 4, Name: "Pineapple", Quantity: 55, Price: 25.00},
	}

	handleRequests()

}
