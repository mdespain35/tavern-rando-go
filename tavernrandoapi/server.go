package main

import (
	//"fmt"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"tavernRando/generator"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {
	generator.PopulateGlobalVars([]string{r.Header["optimized"][0], r.Header["level"][0]})
	if r.URL.Path[1:] == "optimized" {
		generator.Optimized = true
	}
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
	openAIToken := os.Getenv("OPEN_AI_KEY")
	if openAIToken == "" {
		log.Fatal("openAIToken not set")
	}

	client := openai.NewClient(openAIToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo0125,
			MaxTokens: 50,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
	// Check if the call was from the command line or API Call
	if len(os.Args) != 1 {
		generator.PopulateGlobalVars(os.Args[1:])
		fmt.Println(generator.CreatePlayerCharacter())
	}
	// } else {
	// 	http.HandleFunc("/chaos", randomHandler)
	// 	http.HandleFunc("/optimized", randomHandler)
	// 	log.Fatal(http.ListenAndServe(":8080", nil))
	// }
}
