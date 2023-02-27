package main

import (
	"database/sql"
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
		return err
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	return nil
}

// Method #2
// To spin up a new http server
func (app *App) Run(addr string) {
	log.Fatalln(http.ListenAndServe(addr, app.Router))
}
