package main

import (
	"fmt"
	"math/rand"
)

// TODO: Implement the rest of the PlayerCharacter fields
type PlayerCharacter struct {
	// name          string
	race string
	// background    string
	level int
	//class         []PlayerClass
	abilityScores AbilityScores
	// hitpoints     int
}

func (p PlayerCharacter) String() string {
	return fmt.Sprintf("Player Random\nRace: %s\tLevel: %d\n%s", p.race, p.level, p.abilityScores.String())
}

// GenerateAbilityScores generates the ability scores and their modifiers.
// TODO: Make it able to optimize based on player's class/race
func (p *PlayerCharacter) GenerateAbilityScores() {
	if p.abilityScores == nil {
		p.abilityScores = make(map[string]AbilityScore)
	}
	for _, i := range PlayerStats {
		newScore := ability()
		p.abilityScores[i] = AbilityScore{
			score:    newScore,
			modifier: modifier(newScore),
		}
	}
}

func (p *PlayerCharacter) GenerateLevel() {
	p.level = rand.Intn(20) + 1
}

func (p *PlayerCharacter) GenerateRace() {
	p.race = Races[rand.Intn(len(Races))]
}

func main() {
	var player PlayerCharacter

	player.GenerateRace()
	player.GenerateAbilityScores()
	player.GenerateLevel()

	fmt.Println(player)
}
