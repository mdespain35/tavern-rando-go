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
	level         int
	class         []PlayerClass
	AbilityScores AbilityScores
	hitpoints     int
}

var Optimized bool

func (p PlayerCharacter) String() string {
	player := ""
	player += fmt.Sprintf("Player Random\nRace: %s\tLevel: %d\nHitpoints: %d\n", p.race, p.level, p.hitpoints)
	for _, c := range p.class {
		player += c.String()
	}
	player += p.AbilityScores.String()
	return player
}

func (p *PlayerCharacter) GeneratePlayerClass() {
	if p.class == nil {
		p.class = []PlayerClass{}
	}
	p.class = append(p.class, GenerateClass(verifyClass(*p, PickClass())))
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
	p.hitpoints = p.class[0].hitDie
}

func (p *PlayerCharacter) LevelUp() {
	p.level++
	leveledClass := 0

	if rand.Intn(20)+1 == 20 { // Roll a natural 20 to multiclass
		p.GeneratePlayerClass()
		leveledClass = len(p.class) - 1
	} else {
		if len(p.class) > 1 {
			leveledClass = rand.Intn(len(p.class))
		}
		p.class[leveledClass].LevelUp()
		if p.class[leveledClass].level%4 == 0 {
			updateChaosScores(p)
		}
	}
	p.hitpoints += rand.Intn(p.class[leveledClass].hitDie) + 1
}

func verifyClass(p PlayerCharacter, c string) string {
	for _, pc := range p.class {
		if pc.name == c {
			return verifyClass(p, PickClass())
		}
	}
	return c
}

func (p *PlayerCharacter) ChaosScores() {
	for _, i := range PlayerStats {
		newScore := ability()
		p.AbilityScores[i] = AbilityScore{
			Score:    newScore,
			Modifier: modifier(newScore),
		}
	}
	chaosRacialBonus(p)
}

func chaosRacialBonus(p *PlayerCharacter) {
	// Randomized Racial bonus
	for i := 2; i > 0; i-- {
		bonus := rand.Intn(len(PlayerStats))
		ability := p.AbilityScores[PlayerStats[bonus]]
		ability.IncreaseAbilityScore(i)
		ability.UpdateModifier()
		p.AbilityScores[PlayerStats[bonus]] = ability
	}
}

// TODO: Finish implementing this, needs to check to make sure AS doesn't exceed 20
func updateChaosScores(p *PlayerCharacter) {
	points := 2
	for points > 0 {
		bonus := rand.Intn(len(PlayerStats))
		if p.AbilityScores[PlayerStats[bonus]].Score+points < 20 {
			p.AbilityScores[PlayerStats[bonus]] = updateScores(p.AbilityScores[PlayerStats[bonus]], points)
			points -= points
		} else if p.AbilityScores[PlayerStats[bonus]].Score < 20 {
			p.AbilityScores[PlayerStats[bonus]] = updateScores(p.AbilityScores[PlayerStats[bonus]], 1)
			points--
		}
	}
}

func updateScores(as AbilityScore, bonus int) AbilityScore {
	update := as
	update.IncreaseAbilityScore(bonus)
	update.UpdateModifier()
	return update
}

// TODO: Make it able to optimize based on player's class
func (p *PlayerCharacter) OptimizedScores() {

}

func GenerateLevel() int {
	return rand.Intn(20) + 1
}

func (p *PlayerCharacter) GenerateRace() {
	p.race = Races[rand.Intn(len(Races))]
}

func CreatePlayerCharacter() PlayerCharacter {
	var player PlayerCharacter
	player.level = 1
	player.GenerateRace()
	player.GeneratePlayerClass()
	player.GenerateAbilityScores()

	for i := 1; i < GenerateLevel(); i++ {
		player.LevelUp()
	}
	player.hitpoints += player.AbilityScores["Constitution"].Modifier * player.level

	return player
}

func main() {
	Optimized = false
	player := CreatePlayerCharacter()

	fmt.Println(player)
}
