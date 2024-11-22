module tavernRando/generatorAPI

go 1.22.0

replace tavernRando/generator => ./characterGenerator

require (
	github.com/joho/godotenv v1.5.1
	tavernRando/generator v1.1.0
	github.com/rs/cors v1.11.1
)
