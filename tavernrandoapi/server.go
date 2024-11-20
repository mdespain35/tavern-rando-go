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
	"github.com/rs/cors"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {
	params, _ := url.ParseQuery(r.URL.RawQuery)
	optimized, targetLevel := generator.PopulateGlobalVars([]string{params["optimized"][0], params["level"][0]})
	character := generator.CreatePlayerCharacter(optimized, targetLevel)
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
		optimized, targetLevel := generator.PopulateGlobalVars(os.Args[1:])
		fmt.Println(generator.CreatePlayerCharacter(optimized, targetLevel))
	} else {
		mux := http.NewServeMux()
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
			log.Printf("defaulting to port %s", port)
		}
		mux.HandleFunc("/", randomHandler)
		handler := cors.Default().Handler(mux)
		if err := http.ListenAndServe(":"+port, handler); err != nil {
			log.Fatal(err)
		}
	}
}
