package main

import (
	//"fmt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"tavernRando/generator"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {
	generator.PopulateGlobalVars()
	if r.URL.Path[1:] == "optimized" {
		generator.Optimized = true
	}
	character := generator.CreatePlayerCharacter()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(character)
}

func main() {
	// Check if the call was from the command line or API Call
	if len(os.Args) != 1 {
		generator.PopulateGlobalVars()
		fmt.Println(generator.CreatePlayerCharacter())
	} else {
		http.HandleFunc("/chaos", randomHandler)
		http.HandleFunc("/optimized", randomHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
