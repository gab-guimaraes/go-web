package main

import (
	"chapter-4/model"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = "my-secret-pw"
	DBDbase = "csm"
	PORT    = ":8080"
)

var database *sql.DB

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := model.Page{}

	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE page_guid=?",
		pageGUID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)

	fmt.Println(thisPage)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println(err.Error())
	}

	t, _ := template.ParseFiles("templates/blog.html")
	t.Execute(w, thisPage)
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass,
		DBHost, DBDbase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("couldn't connect!")
		log.Println(err.Error())
	}
	database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", ServePage)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)
}
