package main

import (
	"fmt"
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welocme to the homepage!")
	fmt.Println("Endpoint reached: homepage!")
}

func main() {

	http.HandleFunc("/", homepage)             //URL mapper. You can set the URL as "/home", if you want this rout to be the homepage.
	http.ListenAndServe("localhost:8000", nil) //set's up a web-server, here handler is nil

}
