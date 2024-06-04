package main

import (
	//"fmt"
	"encoding/json"
	"log"
	"net/http"
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
	//Stuff for testing JSON Conversion

	// generator.PopulateGlobalVars()
	// player := generator.CreatePlayerCharacter()

	// b, err := json.Marshal(player)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(b))

	http.HandleFunc("/chaos", randomHandler)
	http.HandleFunc("/optimized", randomHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
