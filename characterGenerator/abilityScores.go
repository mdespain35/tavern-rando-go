package characterGenerator

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

type AbilityScore struct {
	Score    int
	Modifier int
}

type AbilityScores map[string]AbilityScore

// Stringer method
func (aS AbilityScores) String() string {
	statBlock := ""
	for _, scores := range PlayerStats {
		statBlock += fmt.Sprintf("%s\n%s", scores, aS[scores].String())
	}
	return statBlock
}

// Stringer method
func (a AbilityScore) String() string {
	return fmt.Sprintf("Score: %d\tModifier: %d\n", a.Score, a.Modifier)
}

// IncreaseAbilityScore takes an int and increases the ability score.
func (a *AbilityScore) IncreaseAbilityScore(increase int) {
	a.Score += increase
}

// UpdateModifier calculates the change in the modifier after an ability score change has occurred.
func (a *AbilityScore) UpdateModifier() {
	a.Modifier = modifier(a.Score)
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
	slices.Reverse(rolls)
	for j := 0; j < 3; j++ {
		sum += rolls[j]
	}
	return sum
}
