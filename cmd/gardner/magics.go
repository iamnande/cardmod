package main

import (
	"context"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/repositories"
)

// offensive magics.
var offensiveMagics = []string{
	"Water",
	"Aero",
	"Bio",
	"Demi",
	"Quake",
	"Tornado",
	"Holy",
	"Flare",
	"Meteor",
	"Ultima",
	"Apocalypse",
	"Fire", "Fira", "Firaga",
	"Blizzard", "Blizzara", "Blizzaga",
	"Thunder", "Thundara", "Thundaga",
}

// restorative magics.
var restorativeMagics = []string{
	"Esuna",
	"Cure", "Cura", "Curaga",
	"Life", "Full-Life", "Regen",
}

// indirect magics.
var indirectMagics = []string{
	"Scan",
	"Sleep",
	"Blind",
	"Silence",
	"Confuse",
	"Berserk",
	"Break",
	"Zombie",
	"Death",
	"Double",
	"Triple",
	"Dispel",
	"Protect",
	"Shell",
	"Reflect",
	"Float",
	"Drain",
	"Haste",
	"Slow", "Stop",
	"Meltdown",
	"Pain",
	"Aura",
}

func seedMagics(ctx context.Context, db *database.Client) error {

	// NOTE: this gets wayyyy easier with known levels in the model
	magics := make([]string, 0)
	magics = append(magics, offensiveMagics...)
	magics = append(magics, restorativeMagics...)
	magics = append(magics, indirectMagics...)

	// initialize client
	magicDB := repositories.NewMagicRepository(db)

	// create the diff
	for i := 0; i < len(magics); i++ {
		magic, err := magicDB.GetMagic(ctx, magics[i])
		if err != nil {
			if !cerrors.IsAPIError(err) {
				return err
			}
		}

		if magic != nil {
			_, err = magicDB.CreateMagic(ctx, magics[i])
			if err != nil {
				return err
			}
		}
	}

	// success
	return nil

}
