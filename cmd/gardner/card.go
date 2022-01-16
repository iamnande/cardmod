package main

// level one cards.
var levelOneCards = []string{
	"Geezard",
	"Funguar",
	"Bite Bug",
	"Red Bat",
	"Gayla",
	"Gesper",
	"Fastitocalon-F",
	"Blood Soul",
	"Caterchipillar",
	"Cockatrice",
}

// level two cards.
var levelTwoCards = []string{
	"Grat",
	"Buel",
	"Mesmerize",
	"Glacial Eye",
	"Belhelmel",
	"Thrustaevis",
	"Anacondaur",
	"Creeps",
	"Grendel",
	"Jelleye",
	"Grand Mantis",
}

// level three cards.
var levelThreeCards = []string{
	"Forbidden",
	"Armadodo",
	"Tri-Face",
	"Fastitocalon",
	"Snow Lion",
	"Ochu",
	"Death Claw",
	"Cactuar",
	"Tonberry",
	"Abyss Worm",
}

// level four cards.
var levelFourCards = []string{
	"Turtapod",
	"Vysage",
	"T-Rexaur",
	"Bomb",
	"Blitz",
	"Wendigo",
	"Torama",
	"Imp",
	"Blue Dragon",
	"Adamantoise",
	"Hexadragon",
}

// level five cards.
var levelFiveCards = []string{
	"Iron Giant",
	"Behemoth",
	"Chimera",
	"PuPu",
	"Elastoid",
	"GIM47N",
	"Malboro",
	"Elnoyle",
	"Tonberry King",
	"Wedge, Biggs",
}

// level six cards.
var levelSixCards = []string{
	"Fujin, Raijin",
	"Elvoret",
	"X-ATM092",
	"Granaldo",
	"Gerogero",
	"Iguion",
	"Abadon",
	"Trauma",
	"Oilboyle",
	"Shumi Tribe",
	"Krysta",
}

// level seven cards.
var levelSevenCards = []string{
	"Propagator",
	"Jumbo Cactuar",
	"Tri-Point",
	"Gargantua",
	"Mobile Type 8",
	"Sphinxara",
	"Tiamat",
	"BGH251F2",
	"Red Giant",
	"Catoblepas",
	"Ultima Weapon",
}

// level eight cards.
var levelEightCards = []string{
	"Chubby Chocobo",
	"Angelo",
	"Gilgamesh",
	"MiniMog",
	"Chicobo",
	"Quezacotl",
	"Shiva",
	"Ifrit",
	"Siren",
	"Sacred",
	"Minotaur",
}

// level nine cards.
var levelNineCards = []string{
	"Carbuncle",
	"Diablos",
	"Leviathan",
	"Odin",
	"Pandemona",
	"Cerberus",
	"Alexander",
	"Phoenix",
	"Bahamut",
	"Doomtrain",
	"Eden",
}

// level ten cards.
var levelTenCards = []string{
	"Ward",
	"Kiros",
	"Laguna",
	"Selphie",
	"Quistis",
	"Irvine",
	"Zell",
	"Rinoa",
	"Edea",
	"Seifer",
	"Squall",
}

// LoadCards is a function to construct a complete list of cards from
// the lists of the various levels of cards.
func LoadCards() []string {

	// complete list of cards
	// TODO: add card level metadata (struct)
	var cards = make([]string, 0)

	// add each level of cards
	cards = append(cards, levelOneCards...)
	cards = append(cards, levelTwoCards...)
	cards = append(cards, levelThreeCards...)
	cards = append(cards, levelFourCards...)
	cards = append(cards, levelFiveCards...)
	cards = append(cards, levelSixCards...)
	cards = append(cards, levelSevenCards...)
	cards = append(cards, levelEightCards...)
	cards = append(cards, levelNineCards...)
	cards = append(cards, levelTenCards...)

	// return complete list of cards
	return cards

}
