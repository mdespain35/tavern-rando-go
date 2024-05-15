package tavernrandogo

import (
	"fmt"
	"math/rand"
	"slices"
)

var PlayerStats = []string{
	"Strength",
	"Dexterity",
	"Constitution",
	"Intelligence",
	"Wisdom",
	"Charisma",
}

type AbilityScores map[string]AbilityScore

// Stringer method
func (aS AbilityScores) String() string {
	statBlock := ""
	for ability, scores := range aS {
		statBlock += fmt.Sprintf("%s\n%s", ability, scores.String())
	}
	return statBlock
}

type AbilityScore struct {
	score    int
	modifier int
}

// Stringer method
func (a AbilityScore) String() string {
	return fmt.Sprintf("Score: %d\nModifier: %d\n", a.score, a.modifier)
}

// modifier calculates the ability score's modifier which determines the bonus/penalty for that ability.
func modifier(score int) int {
	if score >= 10 || score%2 == 0 {
		return (score - 10) / 2
	} else {
		return (score-10)/2 - 1
	}
}

// ability uses randomness to generate the score for an ability.
func ability() int {
	sum := 0
	rolls := []int{}
	for i := 0; i < 4; i++ {
		rolls = append(rolls, rand.Intn(6)+1)
	}
	slices.Sort(rolls)
	for j := 1; j < 4; j++ {
		sum += rolls[j]
	}
	return sum
}
