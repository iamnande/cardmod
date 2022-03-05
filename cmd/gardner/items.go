package main

import (
	"context"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/repositories"
)

// Item represents a item.
// TODO: add purpose attribute.
// TODO: add effect attribute.
// TODO: add location attribute.
type Item struct {
	name string
}

func (c *Item) Name() string {
	return c.name
}

var (
	// TODO: add random items like magazines, weapons mon mars, magical lamp, etc.
	items = []*Item{

		// Restorative
		{name: "Potion"},
		{name: "Potion+"},
		{name: "Hi-Potion"},
		{name: "Hi-Potion+"},
		{name: "X-Potion"},
		{name: "Mega-Potion"},

		// Revival
		{name: "Phoenix Down"},
		{name: "Mega Phoenix"},

		// Forbidden Medicine
		{name: "Elixir"},
		{name: "Megalixir"},

		// Status Recovery
		{name: "Antidote"},
		{name: "Soft"},
		{name: "Eye Drops"},
		{name: "Echo Screen"},
		{name: "Holy Water"},
		{name: "Remedy"},
		{name: "Remedy+"},

		// Invincibility
		{name: "Hero-trial"},
		{name: "Hero"},
		{name: "Holy War-trial"},
		{name: "Holy War"},

		// Spell Stones
		{name: "Shell Stone"},
		{name: "Protect Stone"},
		{name: "Aura Stone"},
		{name: "Death Stone"},
		{name: "Holy Stone"},
		{name: "Flare Stone"},
		{name: "Meteor Stone"},
		{name: "Ultima Stone"},

		// GF Summon
		{name: "Gysahl Greens"},
		{name: "Phoenix Pinion"},
		{name: "Friendship"},

		// Shelters
		{name: "Tent"},
		{name: "Pet House"},
		{name: "Cottage"},

		// GF Recovery
		{name: "G-Potion"},
		{name: "G-Hi-Potion"},
		{name: "G-Mega-Potion"},
		{name: "G-Returner"},

		// Rename Card
		{name: "Rename Card"},

		// GF Ability
		{name: "HP-J Scroll"},
		{name: "Str-J Scroll"},
		{name: "Vit-J Scroll"},
		{name: "Mag-J Scroll"},
		{name: "Spr-J Scroll"},
		{name: "Luck-J Scroll"},
		{name: "Aegis Amulet"},
		{name: "Elem Atk"},
		{name: "Elem Guard"},
		{name: "Status Atk"},
		{name: "Status Guard"},
		{name: "Rosetta Stone"},

		// Command Ability
		{name: "Magic Scroll"},
		{name: "GF Scroll"},
		{name: "Draw Scroll"},
		{name: "Item Scroll"},
		{name: "Gambler Spirit"},
		{name: "Healing Ring"},
		{name: "Phoenix Spirit"},
		{name: "Med Kit"},
		{name: "Bomb Spirit"},
		{name: "Hungry Cookpot"},
		{name: "Meg's Amulet"},

		// GF Enhancement
		{name: "Stell Pipe"},
		{name: "Star Fragment"},
		{name: "Energy Crystal"},
		{name: "Samantha Soul"},
		{name: "Healing Mail"},
		{name: "Silver Sail"},
		{name: "Gold Armor"},
		{name: "Diamond Armor"},

		// Character Ability
		{name: "Regen Ring"},
		{name: "Giant's Ring"},
		{name: "Gaea's Ring"},
		{name: "Strength Love"},
		{name: "Power Wrist"},
		{name: "Hyper Wrist"},
		{name: "Turtle Shell"},
		{name: "Orihalcon"},
		{name: "Adamantine"},
		{name: "Rune Armlet"},
		{name: "Force Armlet"},
		{name: "Magic Armlet"},
		{name: "Circlet"},
		{name: "Hypno Crown"},
		{name: "Royal Crown"},
		{name: "Jet Engine"},
		{name: "Rocket Engine"},
		{name: "Moon Curain"},
		{name: "Steel Curtain"},
		{name: "Glow Curtain"},
		{name: "Accelerator"},
		{name: "Monk's Code"},
		{name: "Knight's Code"},
		{name: "Doc's Code"},
		{name: "Hundred Needles"},
		{name: "Three Stars"},
		{name: "Ribbon"},

		// Ammo
		{name: "Normal Ammo"},
		{name: "Shotgun Ammo"},
		{name: "Dark Ammo"},
		{name: "Fire Ammo"},
		{name: "Demolition Ammo"},
		{name: "Fast Ammo"},
		{name: "AP Ammo"},
		{name: "Pulse Ammo"},

		// Refinement
		{name: "M-Stone Piece"},
		{name: "Magic Stone"},
		{name: "Wizard Stone"},
		{name: "Ochu Tentacle"},
		{name: "Healing Water"},
		{name: "Cockatrice Pinion"},
		{name: "Zombie Powder"},
		{name: "Lightweight"},
		{name: "Sharp Spike"},
		{name: "Screw"},
		{name: "Saw Blade"},
		{name: "Mesmerize Blade"},
		{name: "Vampire Fang"},
		{name: "Fury Fragment"},
		{name: "Betrayal Sword"},
		{name: "Sleep Powder"},
		{name: "Life Ring"},
		{name: "Dragon Fang"},

		// Blue Magic (Quistis LBs)
		{name: "Spider Web"},
		{name: "Coral Fragment"},
		{name: "Curse Spike"},
		{name: "Black Hole"},
		{name: "Water Crystal"},
		{name: "Missile"},
		{name: "Mystery Fluid"},
		{name: "Running Fire"},
		{name: "Inferno Fang"},
		{name: "Malboro Tentacle"},
		{name: "Whisper"},
		{name: "Laser Cannon"},
		{name: "Barrier"},
		{name: "Power Generator"},
		{name: "Dark Matter"},

		// GF Compatibility
		{name: "Bomb Fragment"},
		{name: "Red Fang"},
		{name: "Arctic Wind"},
		{name: "North Wind"},
		{name: "Dynamo Stone"},
		{name: "Shear Feather"},
		{name: "Venom Fang"},
		{name: "Steel Orb"},
		{name: "Moon Stone"},
		{name: "Dino Bone"},
		{name: "Windmill"},
		{name: "Dragon Skin"},
		{name: "Fish Fin"},
		{name: "Dragon Fin"},
		{name: "Silence Powder"},
		{name: "Poison Powder"},
		{name: "Dead Spirit"},
		{name: "Chef's Knife"},
		{name: "Cactus Thorn"},
		{name: "Shaman Stone"},

		// Fuel
		{name: "Fuel"},

		// Stat Boosting
		{name: "HP Up"},
		{name: "Str Up"},
		{name: "Vit Up"},
		{name: "Mag Up"},
		{name: "Spr Up"},
		{name: "Spd Up"},
		{name: "Luck Up"},

		// LuvLuv G
		{name: "LuvLuv G"},
	}
)

func seedItems(ctx context.Context, db *database.Client) error {

	// initialize client
	itemDB := repositories.NewItemRepository(db)

	// create the diff
	for i := 0; i < len(items); i++ {
		item, err := itemDB.GetItem(ctx, items[i].Name())
		if err != nil {
			if !cerrors.IsAPIError(err) {
				return err
			}
		}

		if item == nil {
			_, err = itemDB.CreateItem(ctx, items[i].Name())
			if err != nil {
				return err
			}
		}
	}

	// success
	return nil

}
