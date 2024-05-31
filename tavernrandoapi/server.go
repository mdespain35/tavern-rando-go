package main

import (
	"fmt"
	"tavernRando/generator"
)

func main() {
	generator.PopulateGlobalVars()
	player := generator.CreatePlayerCharacter()

	fmt.Println(player)
}
