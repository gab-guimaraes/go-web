package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	Port = ":8080"
)

const (
	game = "resident evil"
)

func serverDynamic(w http.ResponseWriter, r *http.Request) {
	gameList := [...]string{"mario bros deluxe", "resident evil 7"}
	response := "The time is now " + time.Now().String()
	response = response + "\n" + game
	for _, item := range gameList {
		response += item + "\n"
	}
	fmt.Fprint(w, response)
}

func serverStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}

func main() {
	http.HandleFunc("/static", serverStatic)
	http.HandleFunc("/", serverDynamic)
	http.ListenAndServe(Port, nil)
}
