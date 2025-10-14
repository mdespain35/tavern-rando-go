package characterGenerator

type Skill struct {
	Name string `json:"name"`
	Stat string `json:"stat"`
}

var SkillsList = []Skill{
	{
		Name: "Acrobatics",
		Stat: "Dexterity",
	},
	{
		Name: "Animal Handling",
		Stat: "Wisdom",
	},
	{
		Name: "Arcana",
		Stat: "Intelligence",
	},
	{
		Name: "Athletics",
		Stat: "Strength",
	},
	{
		Name: "Deception",
		Stat: "Charisma",
	},
	{
		Name: "History",
		Stat: "Intelligence",
	}, {
		Name: "Insight",
		Stat: "Wisdom",
	},
	{
		Name: "Intimidation",
		Stat: "Charisma",
	},
	{
		Name: "Investigation",
		Stat: "Intelligence",
	},
	{
		Name: "Medicine",
		Stat: "Wisdom",
	},
	{
		Name: "Nature",
		Stat: "Intelligence",
	},
	{
		Name: "Perception",
		Stat: "Wisdom",
	},
	{
		Name: "Performance",
		Stat: "Charisma",
	},
	{
		Name: "Persuasion",
		Stat: "Charisma",
	},
	{
		Name: "Religion",
		Stat: "Intelligence",
	},
	{
		Name: "Sleight of Hand",
		Stat: "Dexterity",
	},
	{
		Name: "Stealth",
		Stat: "Dexterity",
	},
	{
		Name: "Survival",
		Stat: "Wisdom",
	},
}

var ArtisanTools = []string{
	"Alchemist's Supplies",
	"Brewer's Supplies",
	"Calligrapher's Supplies",
	"Carpenter's Tools",
	"Cartographer's Tools",
	"Cobbler's Tools",
	"Cook's Utensils",
	"Glassblower's Tools",
	"Jewler's Tools",
	"Leatherworker's Tools",
	"Mason's Tools",
	"Painter's Supplies",
	"Potter's Tools",
	"Smith's Tools",
	"Tinker's Tools",
	"Weaver's Tools",
	"Woodcarver's Tools",
}

var GamingSet = []string{
	"Dice Set",
	"Dragonchess Set",
	"Playing Card Set",
	"Three-Dragon Ante Set",
}

var MusicalInstruments = []string{
	"Bagpipes",
	"Drum",
	"Dulcimer",
	"Flute",
	"Lute",
	"Lyre",
	"Horn",
	"Pan Flute",
	"Shawm",
	"Viol",
}

var MiscTools = []string{
	"Water Vehicles",
	"Land Vehicles",
	"Disguise Kit",
	"Forgery Kit",
	"Herbalism Kit",
	"Navigator's Tools",
	"Poisoner's Kit",
	"Thieve's Tools",
}
