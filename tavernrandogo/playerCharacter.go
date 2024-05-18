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
	AbilityScores AbilityScores
	// hitpoints     int
}

var Optimized bool

func (p PlayerCharacter) String() string {
	return fmt.Sprintf("Player Random\nRace: %s\tLevel: %d\n%s", p.race, p.level, p.AbilityScores.String())
}

// GenerateAbilityScores generates the ability scores and their modifiers.
func (p *PlayerCharacter) GenerateAbilityScores() {
	if p.AbilityScores == nil {
		p.AbilityScores = make(map[string]AbilityScore)
	}
	if Optimized {
		p.OptimizedScores()
	} else {
		p.ChaosScores()
	}

}

func (p *PlayerCharacter) ChaosScores() {
	for _, i := range PlayerStats {
		newScore := ability()
		p.AbilityScores[i] = AbilityScore{
			Score:    newScore,
			Modifier: modifier(newScore),
		}
	}
	// Randomized Racial bonus
	for i := 2; i > 0; i-- {
		bonus := rand.Intn(len(PlayerStats))
		ability := p.AbilityScores[PlayerStats[bonus]]
		ability.IncreaseAbilityScore(i)
		ability.UpdateModifier()
		p.AbilityScores[PlayerStats[bonus]] = ability
	}

}

// TODO: Make it able to optimize based on player's class
func (p *PlayerCharacter) OptimizedScores() {

}

func (p *PlayerCharacter) GenerateLevel() {
	p.level = rand.Intn(20) + 1
}

func (p *PlayerCharacter) GenerateRace() {
	p.race = Races[rand.Intn(len(Races))]
}

func main() {
	var player PlayerCharacter
	Optimized = false
	player.GenerateRace()
	player.GenerateAbilityScores()
	player.GenerateLevel()

	fmt.Println(player)
}
