module tavernRando/generatorAPI

go 1.22.0

replace tavernRando/generator => ../characterGenerator

require (
	github.com/joho/godotenv v1.5.1
	github.com/sashabaranov/go-openai v1.25.0
	tavernRando/generator v1.1.0
)
