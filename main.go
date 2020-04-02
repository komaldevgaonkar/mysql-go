package main

import (
	"encoding/json"
	"mysql-go/db"
	"net/http"

	"github.com/gorilla/mux"
)

type employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	http.ListenAndServe(":8000", router)
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := db.GetMySQLDB()
	var employees []employee
	result, err := db.Query("SELECT id, name, city from employee")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post employee
		err := result.Scan(&post.ID, &post.Name, &post.City)
		if err != nil {
			panic(err.Error())
		}
		employees = append(employees, post)
	}
	json.NewEncoder(w).Encode(employees)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	db, err := db.GetMySQLDB()
	result, err := db.Query("SELECT id, name, city FROM employee WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post employee
	for result.Next() {
		err := result.Scan(&post.ID, &post.Name, &post.City)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(post)
}
