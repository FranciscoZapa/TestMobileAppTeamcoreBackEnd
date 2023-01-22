package main

import (
	"encoding/json"
	"net/http"

	"github.com/FranciscoZapa/TestMobileAppTeamcoreBackEnd/src/api"
)

func main() {

	response := api.GetData()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}
