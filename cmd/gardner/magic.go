package main

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

// LoadMagics is a function to construct a complete list of magics from
// the lists of indirect, restorative, and offensive magics.
func LoadMagics() []string {

	// complete list of magics
	// TODO: add magic categories (struct)
	// TODO: add stat/elem/status junction values (ex: Fire - +50% ATK)
	var magics = make([]string, 0)

	// add each type of magic
	magics = append(magics, offensiveMagics...)
	magics = append(magics, restorativeMagics...)
	magics = append(magics, indirectMagics...)

	// return complete list of magics
	return magics

}
