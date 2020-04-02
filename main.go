package main

import (
	"encoding/json"
	"log"
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
	dbconn, err := db.GetMySQLDB()
	if err != nil {
		log.Println("Connection Failed to Open")
	} else {
		log.Println("Connection Established")
	}

	router := mux.NewRouter()
	router.HandleFunc("/posts", getPosts).Methods("GET")
	http.ListenAndServe(":8000", router)
}
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []employee
	result, err := dbconn.Query("SELECT id, name, city from employee")
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
