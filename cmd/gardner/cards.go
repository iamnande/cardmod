package main

import (
	"context"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/repositories"
)

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

func seedCards(ctx context.Context, db *database.Client) error {

	// NOTE: this gets wayyyy easier with known levels in the model
	cards := make([]string, 0)
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

	// initialize client
	cardDB := repositories.NewCardRepository(db)

	// create the diff
	for i := 0; i < len(cards); i++ {
		card, err := cardDB.GetCard(ctx, cards[i])
		if err != nil {
			if !cerrors.IsAPIError(err) {
				return err
			}
		}

		if card != nil {
			_, err = cardDB.CreateCard(ctx, cards[i])
			if err != nil {
				return err
			}
		}
	}

	// success
	return nil

}
