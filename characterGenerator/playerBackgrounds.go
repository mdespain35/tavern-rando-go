package generator

import "fmt"

type Background struct {
	name string
	// TODO: Add in skill proficiencies that come with each background
	// TODO: Add in background feature
}

var BackgroundOptions = []Background{
	{
		name: "Acolyte",
		// Insight, Religion
	},
	{
		name: "Charlatan",
		// Deception, Sleight of Hand
	},
	{
		name: "Criminal",
		// Deception, Stealth
	},
	{
		name: "Entertainer",
		// Acrobatics, Performance
	},
	{
		name: "Fisher",
		// History, Survival
	},
	{
		name: "Folk Hero",
		// Animal Handling, Survival
	},
	{
		name: "Gladiator",
		// Acrobatics, Peformance
	},
	{
		name: "Guild Artisan",
		// Insight, Persuasion
	},
	{
		name: "Hermit",
		// Medicine, Religion
	},
	{
		name: "Knight",
		// History, Persuasion
	},
	{
		name: "Marine",
		// Athletics, Survival
	},
	{
		name: "Noble",
		// History, Persuasion
	},
	{
		name: "Outlander",
		// Athletics, Survival
	},
	{
		name: "Pirate",
		// Athletics, Perception
	},
	{
		name: "Sage",
		// Arcana, History
	},
	{
		name: "Sailor",
		// Athletics, Perception
	},
	{
		name: "Shipwright",
		// History, Perception
	},
	{
		name: "Smuggler",
		// Athletics, Deception
	},
	{
		name: "Soldier",
		// Athletics, Intimidation
	},
	{
		name: "Urchin",
		// Sleight of Hand, Stealth
	},
}

// Stringer method
func (b Background) String() string {
	return fmt.Sprintf("Background: %s", b.name)
}
