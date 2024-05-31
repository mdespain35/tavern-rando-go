package generator

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
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
var TargetLevel int

func (p PlayerCharacter) String() string {
	player := ""
	player += fmt.Sprintf("Player Random\nRace: %s\tLevel: %d\nHitpoints: %d\n", p.race, p.level, p.hitpoints)
	for _, c := range p.class {
		player += c.String()
	}
	player += p.AbilityScores.String()
	return player
}

// GeneratePlayerClass initializes the p.class field if needed and then adds a PlayerClass to the slice
func (p *PlayerCharacter) GeneratePlayerClass() {
	if p.class == nil {
		p.class = []PlayerClass{}
	}
	p.class = append(p.class, GenerateClass(verifyClass(*p, PickClass())))
}

// GenerateAbilityScores generates the ability scores and their modifiers, should only be called once.
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

// LevelUp levels up the Player's character, while also handling the leveling for the Player's class by using
// the PlayerClass definition of LevelUp.
func (p *PlayerCharacter) LevelUp() {
	p.level++
	leveledClass := 0
	upScores := false // Flag marking if ability score improvements needed.

	// TODO: Add a check to ensure the character has the stats to multiclass otherwise pick a different class.
	if rand.Intn(20)+1 == 20 { // Roll a natural 20 to multiclass
		p.GeneratePlayerClass()
		leveledClass = len(p.class) - 1
	} else {
		if len(p.class) > 1 {
			leveledClass = rand.Intn(len(p.class))
		}
		p.class[leveledClass].LevelUp()
		// Block checking if an ability score improvement is needed.
		if p.class[leveledClass].level%4 == 0 || p.class[leveledClass].level == 19 {
			if p.class[leveledClass].name == "Barbarian" && p.class[leveledClass].level == 20 {
				p.AbilityScores["Strength"] = updateScores(p.AbilityScores["Strength"], 4)
				p.AbilityScores["Constitution"] = updateScores(p.AbilityScores["Constitution"], 4)
			} else {
				upScores = true
			}
		} else if p.class[leveledClass].name == "Fighter" {
			if p.class[leveledClass].level == 6 || p.class[leveledClass].level == 14 {
				upScores = true
			}
		} else if p.class[leveledClass].name == "Rogue" && p.class[leveledClass].level == 10 {
			upScores = true
		}
	}
	if upScores {
		if Optimized {
			updateOptmizeScores(p, leveledClass)
		} else {
			updateChaosScores(p, 2)
		}
	}
	p.hitpoints += rand.Intn(p.class[leveledClass].hitDie) + 1 // Health increases based on which class was leveled up.
}

// verifyClass ensures that when multiclassing, an already selected class is not added.
func verifyClass(p PlayerCharacter, c string) string {
	for _, pc := range p.class {
		if pc.name == c {
			return verifyClass(p, PickClass())
		}
	}
	return c
}

// ChaosScores creates the ability scores for unoptimized characters for a truly random experience.
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

// chaosRacialBonus adds a random racial bonus to 2 ability scores rather than using the preferred stats of the Player's class.
func chaosRacialBonus(p *PlayerCharacter) {
	// Randomized Racial bonus
	for i := 2; i > 0; i-- {
		bonus := rand.Intn(len(PlayerStats))
		p.AbilityScores[PlayerStats[bonus]] = updateScores(p.AbilityScores[PlayerStats[bonus]], i)
	}
}

// updateChaosScores updates the Player's ability scores randomly when the Player's class hits a level that triggers
// an ability score improvement.
func updateChaosScores(p *PlayerCharacter, point int) {
	points := point
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

// updateScores handles the actual updating of the ability scores.
func updateScores(as AbilityScore, bonus int) AbilityScore {
	update := as
	update.IncreaseAbilityScore(bonus)
	update.UpdateModifier()
	return update
}

// OptimizedScores generates a character that applies the highest rolled stats to the preferred stats of the character's class.
func (p *PlayerCharacter) OptimizedScores() {
	scores := []int{}
	for i := 0; i < 6; i++ { // Generate ability scores for sorting.
		scores = append(scores, ability())
	}
	slices.Sort(scores)
	slices.Reverse(scores)
	for i := 0; i < len(p.class[0].preferredStats); i++ { // Assigning highest rolls to preferred stats for level 1 class.
		p.AbilityScores[p.class[0].preferredStats[i]] = AbilityScore{
			Score:    scores[i],
			Modifier: modifier(scores[i]),
		}
	}

	for j := len(p.class[0].preferredStats); j < len(scores); j++ {
		for _, a := range PlayerStats {
			if _, ok := p.AbilityScores[a]; !ok { // Fill the ability score if it has not been populated yet.
				p.AbilityScores[a] = AbilityScore{
					Score:    scores[j],
					Modifier: modifier(scores[j]),
				}
				break
			}
		}
	}
	optimizeRacialBonus(p)
}

// optimizedRacialBonus assigns the racial bonus to the preferred stats of the Character's class.
func optimizeRacialBonus(p *PlayerCharacter) {
	p.AbilityScores[p.class[0].preferredStats[0]] = updateScores(p.AbilityScores[p.class[0].preferredStats[0]], 2)
	p.AbilityScores[p.class[0].preferredStats[1]] = updateScores(p.AbilityScores[p.class[0].preferredStats[1]], 1)
}

// updateOptimizedScores prioritizes the preferred stats of the leveled class for the ability score improvement.
func updateOptmizeScores(p *PlayerCharacter, leveledClass int) {
	points := 2
	for _, a := range p.class[leveledClass].preferredStats {
		if p.AbilityScores[a].Score+points < 20 {
			p.AbilityScores[a] = updateScores(p.AbilityScores[a], points)
			points -= points
		} else if p.AbilityScores[a].Score < 20 {
			p.AbilityScores[a] = updateScores(p.AbilityScores[a], 1)
			points--
		}
		if points == 0 { // Break out of loop as soon as bonus is awarded.
			break
		}
	}

	// TODO: Add a feat counter for if the preferred stats are maxed out?
	if points > 0 { // Make sure no points get wasted.
		updateChaosScores(p, points)
	}
}

// GenerateLevel is a helper function for if a level is not specified by a user. May remove later.
func GenerateLevel() int {
	return rand.Intn(20) + 1
}

// GenerateRace populates the p.race field by randomly selecting a race from the Races slice.
func (p *PlayerCharacter) GenerateRace() {
	p.race = Races[rand.Intn(len(Races))]
}

// CreatePlayerCharacter puts all of the pieces of the sandwich together and returns a PlayerCharacter object.
func CreatePlayerCharacter() PlayerCharacter {
	var player PlayerCharacter
	player.level = 1
	player.GenerateRace()
	player.GeneratePlayerClass()
	player.GenerateAbilityScores()

	for i := 1; i < TargetLevel; i++ {
		player.LevelUp()
	}
	// Calculate Constitution bonus at the end in case Constitution was increased during leveling.
	player.hitpoints += player.AbilityScores["Constitution"].Modifier * player.level

	return player
}

// populateGlobalVars is a helper function that reads in input from the CLI and assigns the useful args to the global vars of this program.
func PopulateGlobalVars() {
	// Assign default values in case call is missing one or more variables or bad args are passed.
	Optimized = false
	TargetLevel = GenerateLevel()

	for i := 1; i < len(os.Args); i++ {
		if val, err := strconv.Atoi(os.Args[i]); err == nil { // Check if arg is an int.
			if val > 0 && val <= 20 {
				TargetLevel = val
			}
		} else if boolVal, err := strconv.ParseBool(os.Args[i]); err == nil { // Check if arg is a bool.
			Optimized = boolVal
		}
	}
}
