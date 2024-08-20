package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	APP_PORT = ":4000"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{
		Id:   1,
		Name: "dimas",
	}, {
		Id:   2,
		Name: "xx",
	},
}

func main(){
	http.HandleFunc("/users/get", getUser)

	http.HandleFunc("/users/add", addUser)
	
	log.Println("server running at port", APP_PORT)
	http.ListenAndServe(APP_PORT, nil)

}

func getUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "method not allowed",
		})
	}

	json.NewEncoder(w).Encode(users)
}


func addUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")


	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "method not allowed",
		})
	}

	var req = User{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "method not allowed",
			})
		}
	}

	// generate id

	req.Id = len(users) + 1

	users = append(users, req)


	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)
}
