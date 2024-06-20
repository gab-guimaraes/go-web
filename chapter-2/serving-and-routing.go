package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	_, err := os.Stat(fileName)

	//error? show error.html
	if err != nil {
		fileName = "error.html"
	}
	http.ServeFile(w, r, fileName)
}

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	fmt.Println("welcome" + name)
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	route.HandleFunc("/name/{name}", nameHandler)
	http.Handle("/", route)
	http.ListenAndServe(PORT, nil)
}
