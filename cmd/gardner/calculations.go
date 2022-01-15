package main

type calculation struct {
	card       string
	magic      string
	cardRatio  int32
	magicRatio int32
}

var calculations = []*calculation{
	{
		card:       "Abyss Worm",
		magic:      "Tornado",
		cardRatio:  1,
		magicRatio: 20,
	},
	{
		card:       "Ruby Dragon",
		magic:      "Flare",
		cardRatio:  10,
		magicRatio: 20,
	},
	{
		card:       "Snow Lion",
		magic:      "Blizzaga",
		cardRatio:  1,
		magicRatio: 20,
	},
	{
		card:       "Blitz",
		magic:      "Thundaga",
		cardRatio:  1,
		magicRatio: 20,
	},
	{
		card:       "Hexadragon",
		magic:      "Firaga",
		cardRatio:  3,
		magicRatio: 20,
	},
	{
		card:       "Fastitocalon",
		magic:      "Water",
		cardRatio:  1,
		magicRatio: 50,
	},
	{
		card:       "Thrustaevis",
		magic:      "Aero",
		cardRatio:  1,
		magicRatio: 20,
	},
	{
		card:       "Bomb",
		magic:      "Fira",
		cardRatio:  1,
		magicRatio: 20,
	},
	{
		card:       "Glacial Eye",
		magic:      "Blizzara",
		cardRatio:  1,
		magicRatio: 20,
	},
	{
		card:       "Creeps",
		magic:      "Thundara",
		cardRatio:  1,
		magicRatio: 20,
	},
}
