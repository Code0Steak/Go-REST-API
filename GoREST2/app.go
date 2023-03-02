package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router //router information
	DB     *sql.DB
}

// Method #1: init
// Initializes connection with MySQL Database
func (app *App) Initialize(DB_Username, DB_Pass, DB_Name string) error {
	connectionString := fmt.Sprintf("%v:%v@tcp()/%v", DB_Username, DB_Pass, DB_Name)
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes() //remember to call the handleRoutes method.
	return nil
}

// Method #2
// To spin up a new http server
func (app *App) Run(addr string) {
	log.Fatalln(http.ListenAndServe(addr, app.Router))
}

// This func sends a response to client
func sendResponse(w http.ResponseWriter, statCode int, payload interface{}) {
	response, _ := json.Marshal(payload) //send a JSON response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statCode)
	w.Write(response)
}

// This func will handle all NON-200 responses
func sendError(w http.ResponseWriter, statCode int, err string) {
	err_message := map[string]string{"error": err}
	sendResponse(w, statCode, err_message)
}

// Handle Method #1
// Will handle the 'GET' request for retreiving products
func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: products")
	products, err := getProductsFromDB(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, products)
}

// Handle Method #2
// Will handle the 'GET' request for retreiving a product based on its id
func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: product/{id}")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"]) //vars is a map, retreive the value corresponding to key "id"
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid product ID.") //Status code: 400
		return
	}
	product := product{ID: id}
	err = product.getProductFromDB(app.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows: //if no rows exist for the ID
			sendError(w, http.StatusNotFound, "Product NOT found!")
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendResponse(w, http.StatusOK, product)

}

// Handle Method #3
// Will handle the POST request. The Handler will add a new row to the products table with the data for new product.
func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: /product")

	var p product
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = p.createProductinDB(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusCreated, p)

}

// Handle Method #4
func (app *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: /product/{id} for updating product by id.")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	var p product
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	p.ID = id
	err = p.updateProduct(app.DB)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, p)

}

// Handle Method #5
func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request) {

	log.Println("Endpoint hit: /product/{id} for deleting product by id")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Incorrect id")
	}

	var p product
	p.ID = id
	err = p.deleteProduct(app.DB)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, map[string]string{"result": fmt.Sprintf("Product with ID %v was deleted succesfully.", p.ID)})

}

// Method #3
// Handle all routes
func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/product/{id}", app.getProduct).Methods("GET")
	app.Router.HandleFunc("/product", app.createProduct).Methods("POST")
	app.Router.HandleFunc("/product/{id}", app.updateProduct).Methods("PUT")
	app.Router.HandleFunc("/product/{id}", app.deleteProduct).Methods("DELETE")
}
