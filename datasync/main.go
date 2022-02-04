package datasync

import (
	// "casino_royal/vault/client"
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Status{Status: "Service is running!!!"})
}

func standings(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	// leagueId := params["leagueId"]
	// id, err := strconv.Atoi(leagueId)

	// if err != nil {
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	// httpClient := client.NewHttpClient()
	// result, err := httpClient.League(id)

	// if err != nil {
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	// json.NewEncoder(w).Encode(result)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", health).Methods("GET")
	router.HandleFunc("/standings/{leagueId}", standings).Methods("GET")

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
