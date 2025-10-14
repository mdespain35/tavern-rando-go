package characterGenerator

import (
	"fmt"
	"math/rand"
)

var ClassOptions = []string{
	"Artificer",
	"Barbarian",
	"Bard",
	"Cleric",
	"Druid",
	"Fighter",
	"Monk",
	"Paladin",
	"Ranger",
	"Rogue",
	"Sorcerer",
	"Warlock",
	"Wizard",
	"Blood Hunter",
}

type PlayerClass struct {
	Name            string `json:"name"`
	SubClass        string `json:"subClass"`
	Level           int    `json:"level"`
	HitDie          int    `json:"hitDie"`
	subClassLevel   int
	preferredStats  []string
	subclassOptions []string
}

// LevelUp levels up the PlayerClass and assigns a subclass if the appropriate level is reached.
func (pc *PlayerClass) LevelUp() {
	pc.Level++
	if pc.Level == pc.subClassLevel {
		pc.SubClass = chooseSubclass(pc.subclassOptions)
	}
}

func (pc PlayerClass) String() string {
	return fmt.Sprintf("%s,%s\t%d\n", pc.Name, pc.SubClass, pc.Level)
}

// chooseSubclass returns a random subclass from the PlayerClass's subclassOptions field.
func chooseSubclass(options []string) string {
	return options[rand.Intn(len(options))]
}

// PickClass selects a class randomly from the ClassOptions slice.
func PickClass() string {
	return ClassOptions[rand.Intn(len(ClassOptions))]
}

// GenerateClass populates a PlayerClass object based on a string and then returns the populated object.
func GenerateClass(class string) PlayerClass {
	p := PlayerClass{
		Name:     class,
		SubClass: "",
		Level:    1,
	}
	switch p.Name {
	case "Artificer":
		p.subClassLevel = 3
		p.HitDie = 8
		p.preferredStats = []string{
			"Intelligence", "Dexterity",
		}
		p.subclassOptions = []string{
			"Alchemist", "Armorer", "Artillerist", "Battle Smith",
		}
	case "Barbarian":
		p.subClassLevel = 3
		p.HitDie = 12
		p.preferredStats = []string{
			"Strength", "Constitution",
		}
		p.subclassOptions = []string{
			"Berserker",
			"Totem Warrior",
			"Ancestral Guardian",
			"Storm Herald",
			"Zealot",
			"Beast",
			"Wild Magic",
		}
	case "Bard":
		p.subClassLevel = 3
		p.HitDie = 8
		p.preferredStats = []string{
			"Charisma", "Constitution", "Dexterity",
		}
		p.subclassOptions = []string{
			"College of Lore",
			"College of Valor",
			"College of Glamour",
			"College of Swords",
			"College of Whispers",
			"College of Creation",
			"College of Eloquence",
		}
	case "Blood Hunter":
		p.subClassLevel = 3
		p.HitDie = 10
		if rand.Intn(2) == 0 {
			p.preferredStats = []string{
				"Strength", "Intelligence",
			}
		} else {
			p.preferredStats = []string{
				"Dexterity", "Intelligence",
			}
		}
		p.subclassOptions = []string{
			"Order of the Ghostslayer",
			"Order of the Lycan",
			"Order of the Mutant",
			"Order of the Profane Soul",
		}
	case "Cleric":
		p.subClassLevel = 1
		p.HitDie = 8
		p.preferredStats = []string{
			"Wisdom", "Strength", "Constitution",
		}
		p.subclassOptions = []string{
			"Knowledge Domain",
			"Life Domain",
			"Light Domain",
			"Nature Domain",
			"Tempest Domain",
			"Trickery Domain",
			"War Domain",
			"Forge Domain",
			"Grave Domain",
			"Order Domain",
			"Peace Domain",
			"Twilight Domain",
		}
	case "Druid":
		p.subClassLevel = 2
		p.HitDie = 8
		p.preferredStats = []string{
			"Wisdom", "Constitution",
		}
		p.subclassOptions = []string{
			"Circle of the Land",
			"Circle of the Moon",
			"Circle of Dreams",
			"Circle of the Shepherd",
			"Circle of Spores",
			"Circle of Stars",
			"Circle of Wildfire",
		}
	case "Fighter":
		p.subClassLevel = 3
		p.HitDie = 10
		if rand.Intn(2) == 0 {
			p.preferredStats = []string{
				"Strength", "Constitution",
			}
		} else {
			p.preferredStats = []string{
				"Dexterity", "Constitution",
			}
		}
		p.subclassOptions = []string{
			"Champion",
			"Battle Master",
			"Eldritch Knight",
			"Arcane Archer",
			"Cavalier",
			"Samurai",
			"Psi Warrior",
			"Rune Knight",
		}
	case "Monk":
		p.subClassLevel = 3
		p.HitDie = 8
		p.preferredStats = []string{
			"Dexterity", "Wisdom",
		}
		p.subclassOptions = []string{
			"Way of the Open Hand",
			"Way of Shadow",
			"Way of the Four Elements",
			"Way of the Drunken Master",
			"Way of the Kensei",
			"Way of the Sun Soul",
			"Way of Mercy",
			"Way of the Astral Self",
		}
	case "Paladin":
		p.subClassLevel = 3
		p.HitDie = 10
		p.preferredStats = []string{
			"Strength", "Charisma",
		}
		p.subclassOptions = []string{
			"Oath of Devotion",
			"Oath of the Ancients",
			"Oath of Vengeance",
			"Oath of Conquest",
			"Oath of Redemption",
			"Oath of Glory",
			"Oath of the Watchers",
		}
	case "Ranger":
		p.subClassLevel = 3
		p.HitDie = 10
		p.preferredStats = []string{
			"Dexterity", "Wisdom",
		}
		p.subclassOptions = []string{
			"Hunter",
			"Beast Master",
			"Gloom Stalker",
			"Horizon Walker",
			"Monster Slayer",
			"Fey Wanderer",
			"Swarm Keeper",
		}
	case "Rogue":
		p.subClassLevel = 3
		p.HitDie = 8
		if rand.Intn(2) == 0 {
			p.preferredStats = []string{
				"Dexterity", "Intelligence",
			}
		} else {
			p.preferredStats = []string{
				"Dexterity", "Charisma",
			}
		}
		p.subclassOptions = []string{
			"Thief",
			"Assassin",
			"Arcane Trickster",
			"Inquisitive",
			"Mastermind",
			"Scout",
			"Swashbuckler",
			"Phantom",
			"Soulknife",
		}
	case "Sorcerer":
		p.subClassLevel = 1
		p.HitDie = 6
		p.preferredStats = []string{
			"Charisma", "Constitution",
		}
		p.subclassOptions = []string{
			"Draconic Bloodline",
			"Wild Magic",
			"Divine Soul",
			"Shadow Magic",
			"Storm Sorcery",
			"Aberrant Mind",
			"Clockwork Soul",
		}
	case "Warlock":
		p.subClassLevel = 1
		p.HitDie = 8
		p.preferredStats = []string{
			"Charisma", "Constitution",
		}
		p.subclassOptions = []string{
			"The Archfey",
			"The Fiend",
			"The Great Old One",
			"The Celestial",
			"The Hexblade",
			"The Fathomless",
			"The Genie",
		}
	case "Wizard":
		p.subClassLevel = 2
		p.HitDie = 6
		if rand.Intn(2) == 0 {
			p.preferredStats = []string{
				"Intelligence", "Dexterity",
			}
		} else {
			p.preferredStats = []string{
				"Intelligence", "Constitution",
			}
		}
		p.subclassOptions = []string{
			"School of Abjuration",
			"School of Conjuration",
			"School of Divination",
			"School of Enchantment",
			"School of Evocation",
			"School of Illusion",
			"School of Necromancy",
			"School of Transmutation",
			"War Mage",
			"Bladesinger",
			"Order of Scribes",
		}
	}

	if p.subClassLevel == 1 {
		p.SubClass = chooseSubclass(p.subclassOptions)
	}

	return p
}
