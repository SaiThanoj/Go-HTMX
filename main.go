package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func initDB() {

	var err error

	db, err = sql.Open("mysql", "root:demoapp@(127.0.0.1:3333)/mydb?parseTime=true")

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
	}
}
func main() {

	initDB()
	defer db.Close()

	gRouter := mux.NewRouter()

	gRouter.HandleFunc("/", handler)

	http.ListenAndServe(":3000", gRouter)
}

func handler(w http.ResponseWriter, r *http.Request) {

	err := tmpl.ExecuteTemplate(w, "home.html", nil)

	if err != nil {
		log.Fatal("Error executing template:", err)
	}
}
