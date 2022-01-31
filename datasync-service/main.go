package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//Status of the response
type Status struct {
	Status string `json:"status,omitempty"`
}

func health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Status{Status: "Service is running!!!"})
}

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", health).Methods("GET")

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}

}