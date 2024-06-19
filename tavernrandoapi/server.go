package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"tavernRando/generator"

	"github.com/joho/godotenv"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {
	params, _ := url.ParseQuery(r.URL.RawQuery)
	generator.PopulateGlobalVars([]string{params["optimized"][0], params["level"][0]})
	character := generator.CreatePlayerCharacter()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(character)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Check if the call was from the command line or API Call
	if len(os.Args) != 1 {
		generator.PopulateGlobalVars(os.Args[1:])
		fmt.Println(generator.CreatePlayerCharacter())
	} else {
		http.HandleFunc("/", randomHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
