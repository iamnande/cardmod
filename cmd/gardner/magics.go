package main

import (
	"context"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/repositories"
)

// Magic represents a magic.
type Magic struct {
	name    string
	purpose string // offensive, restorative, indirect
}

func (m *Magic) Name() string {
	return m.name
}

func (m *Magic) Purpose() string {
	return m.purpose
}

var (
	magics = []*Magic{
		{name: "Water", purpose: "Offensive"},
		{name: "Aero", purpose: "Offensive"},
		{name: "Bio", purpose: "Offensive"},
		{name: "Demi", purpose: "Offensive"},
		{name: "Quake", purpose: "Offensive"},
		{name: "Tornado", purpose: "Offensive"},
		{name: "Holy", purpose: "Offensive"},
		{name: "Flare", purpose: "Offensive"},
		{name: "Meteor", purpose: "Offensive"},
		{name: "Ultima", purpose: "Offensive"},
		{name: "Apocalypse", purpose: "Offensive"},
		{name: "Fire", purpose: "Offensive"},
		{name: "Fira", purpose: "Offensive"},
		{name: "Firaga", purpose: "Offensive"},
		{name: "Blizzard", purpose: "Offensive"},
		{name: "Blizzara", purpose: "Offensive"},
		{name: "Blizzaga", purpose: "Offensive"},
		{name: "Thunder", purpose: "Offensive"},
		{name: "Thundara", purpose: "Offensive"},
		{name: "Thundaga", purpose: "Offensive"},
		{name: "Esuna", purpose: "Restorative"},
		{name: "Cure", purpose: "Restorative"},
		{name: "Cura", purpose: "Restorative"},
		{name: "Curaga", purpose: "Restorative"},
		{name: "Life", purpose: "Restorative"},
		{name: "Full-Life", purpose: "Restorative"},
		{name: "Regen", purpose: "Restorative"},
		{name: "Scan", purpose: "Indirect"},
		{name: "Sleep", purpose: "Indirect"},
		{name: "Blind", purpose: "Indirect"},
		{name: "Silence", purpose: "Indirect"},
		{name: "Confuse", purpose: "Indirect"},
		{name: "Berserk", purpose: "Indirect"},
		{name: "Break", purpose: "Indirect"},
		{name: "Zombie", purpose: "Indirect"},
		{name: "Death", purpose: "Indirect"},
		{name: "Double", purpose: "Indirect"},
		{name: "Triple", purpose: "Indirect"},
		{name: "Dispel", purpose: "Indirect"},
		{name: "Protect", purpose: "Indirect"},
		{name: "Shell", purpose: "Indirect"},
		{name: "Reflect", purpose: "Indirect"},
		{name: "Float", purpose: "Indirect"},
		{name: "Drain", purpose: "Indirect"},
		{name: "Haste", purpose: "Indirect"},
		{name: "Slow", purpose: "Indirect"},
		{name: "Stop", purpose: "Indirect"},
		{name: "Meltdown", purpose: "Indirect"},
		{name: "Pain", purpose: "Indirect"},
		{name: "Aura", purpose: "Indirect"},
	}
)

func seedMagics(ctx context.Context, db *database.Client) error {

	// initialize client
	magicDB := repositories.NewMagicRepository(db)

	// create the diff
	for i := 0; i < len(magics); i++ {
		magic, err := magicDB.GetMagic(ctx, magics[i].Name())
		if err != nil {
			if !cerrors.IsAPIError(err) {
				return err
			}
		}

		if magic == nil {
			_, err = magicDB.CreateMagic(ctx, magics[i].Name(), magics[i].Purpose())
			if err != nil {
				return err
			}
		}
	}

	// success
	return nil

}
