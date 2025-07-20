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

type Todo struct {
	Id       int
	TaskName string
	Done     bool
}

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

	gRouter.HandleFunc("/todoList", getTodoList).Methods("GET")

	http.ListenAndServe(":3000", gRouter)
}

func handler(w http.ResponseWriter, r *http.Request) {

	err := tmpl.ExecuteTemplate(w, "home.html", nil)

	if err != nil {
		log.Fatal("Error executing template:", err)
	}
}

func getTodoList(w http.ResponseWriter, r *http.Request) {

	tasks, err := getList()

	if err != nil {
		log.Fatal("Error getting todo list:", err)
	}

	err = tmpl.ExecuteTemplate(w, "todoList", tasks)

	if err != nil {
		log.Fatal("Error executing template:", err)
	}
}

func getList() ([]Todo, error) {

	rows, err := db.Query("SELECT id, taskname, done FROM tasks")

	if err != nil {
		log.Fatal("Error querying the database:", err)
	}
	defer rows.Close()

	var tasks []Todo

	for rows.Next() {
		var task Todo
		if err := rows.Scan(&task.Id, &task.TaskName, &task.Done); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err != nil {
		log.Fatal("Error executing template:", err)
	}

	return tasks, nil
}
