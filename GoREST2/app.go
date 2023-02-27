package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type App struct {
	Router *mux.Router //router information
	DB     *sql.DB
}

// Method #1: init
// Initializes connection with MySQL Database
func (app *App) Initialize() error {
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

// Method #3
// Handle all routes
func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
}
