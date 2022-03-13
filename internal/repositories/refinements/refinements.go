package refinements

var (
	// refinement is the collection of refinement entities.
	refinements = [87]*refinement{
		// Blue Magic
		{source: "Gesper", target: "Black Hole", numerator: 1, denominator: 1},
		{source: "Black Hole", target: "Degenerator", numerator: 1, denominator: 1},
		{source: "Malboro", target: "Malboro Tentacle", numerator: 4, denominator: 1},
		{source: "Malboro Tentacle", target: "Bad Breath", numerator: 1, denominator: 1},
		{source: "Ruby Dragon", target: "Inferno Fang", numerator: 10, denominator: 1},
		{source: "Inferno Fang", target: "Fire Breath", numerator: 1, denominator: 1},
		{source: "Behemoth", target: "Barrier", numerator: 10, denominator: 1},
		{source: "Barrier", target: "Mighty Guard", numerator: 1, denominator: 1},
		{source: "SAM08G", target: "Running Fire", numerator: 1, denominator: 1},
		{source: "Running Fire", target: "Gatling Gun", numerator: 1, denominator: 1},
		{source: "Fastilocalon", target: "Water Crystal", numerator: 1, denominator: 1},
		{source: "Water Crystal", target: "Aqua Breath", numerator: 1, denominator: 1},
		{source: "Creeps", target: "Coral Fragment", numerator: 1, denominator: 1},
		{source: "Coral Fragment", target: "Electrocute", numerator: 1, denominator: 1},
		{source: "Gayla", target: "Mystery Fluid", numerator: 1, denominator: 1},
		{source: "Mystery Fluid", target: "Acid", numerator: 1, denominator: 1},
		{source: "Tri-Face", target: "Curse Spike", numerator: 1, denominator: 1},
		{source: "Curse Spike", target: "L?Death", numerator: 1, denominator: 1},
		{source: "Caterchipillar", target: "Spider Web", numerator: 1, denominator: 1},
		{source: "Spider Web", target: "Ultra Waves", numerator: 1, denominator: 1},

		// Refinement Moment #1
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/preparing-for-the-exam
		{source: "Abyss Worm", target: "Windmill", numerator: 1, denominator: 1},
		{source: "Windmill", target: "Tornado", numerator: 1, denominator: 20},
		{source: "Inferno Fang", target: "Flare", numerator: 1, denominator: 20},
		{source: "Snow Lion", target: "North Wind", numerator: 1, denominator: 1},
		{source: "North Wind", target: "Blizzaga", numerator: 1, denominator: 20},
		{source: "Blitz", target: "Dynamo Stone", numerator: 1, denominator: 1},
		{source: "Dynamo Stone", target: "Thundaga", numerator: 1, denominator: 20},
		{source: "Hexadrago", target: "Red Fang", numerator: 3, denominator: 1},
		{source: "Red Fang", target: "Firaga", numerator: 1, denominator: 20},
		{source: "Fastitocalon-F", target: "Water Crystal", numerator: 5, denominator: 1},
		{source: "Water Crystal", target: "Water", numerator: 1, denominator: 50},
		{source: "Thrustaevis", target: "Shear Feather", numerator: 1, denominator: 1},
		{source: "Shear Feather", target: "Aero", numerator: 1, denominator: 20},
		{source: "Bomb", target: "Bomb Fragment", numerator: 1, denominator: 1},
		{source: "Bomb Fragment", target: "Fira", numerator: 1, denominator: 20},
		{source: "Glacial Eye", target: "Arctic Wind", numerator: 1, denominator: 1},
		{source: "Arctic Wind", target: "Blizzara", numerator: 1, denominator: 20},
		{source: "Coral Fragment", target: "Thundara", numerator: 1, denominator: 20},

		// Refinement Moment #2
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/the-timber-mission
		{source: "Quistis", target: "Samantha Soul", numerator: 1, denominator: 3},
		{source: "Samantha Soul", target: "Triple", numerator: 1, denominator: 60},
		{source: "Zell", target: "Hyper Wrist", numerator: 1, denominator: 3},
		{source: "Diablos", target: "Black Hole", numerator: 1, denominator: 100},
		{source: "Black Hole", target: "Demi", numerator: 1, denominator: 30},
		{source: "Mystery Fluid", target: "Meltdown", numerator: 1, denominator: 10},
		{source: "Curse Spike", target: "Pain", numerator: 1, denominator: 10},
		{source: "Armadodo", target: "Dino Bone", numerator: 1, denominator: 1},
		{source: "T-Rexaur", target: "Dino Bone", numerator: 2, denominator: 1},
		{source: "Dino Bone", target: "Quake", numerator: 1, denominator: 20},
		{source: "Mesmerize", target: "Mesmerize Blade", numerator: 1, denominator: 1},
		{source: "Mesmerize Blade", target: "Regen", numerator: 1, denominator: 20},
		{source: "Tonberry", target: "Chef's Knife", numerator: 1, denominator: 1},
		{source: "Tonberry King", target: "Chef's Knife", numerator: 1, denominator: 1},
		{source: "Chef's Knife", target: "Death", numerator: 1, denominator: 30},
		{source: "Belhelmel", target: "Saw Blade", numerator: 1, denominator: 1},
		{source: "Saw Blade", target: "Death", numerator: 1, denominator: 10},
		{source: "Tent", target: "Curaga", numerator: 1, denominator: 10},
		{source: "Cockatrice", target: "Cockatrice Pinion", numerator: 1, denominator: 1},
		{source: "Cockatrice Pinion", target: "Break", numerator: 1, denominator: 20},
		{source: "Forbidden", target: "Betrayal Sword", numerator: 1, denominator: 1},
		{source: "Betrayal Sword", target: "Confuse", numerator: 1, denominator: 20},
		{source: "Anacondaur", target: "Venom Fang", numerator: 1, denominator: 1},
		{source: "Venom Fang", target: "Bio", numerator: 1, denominator: 20},
		{source: "Malboro Tentacle", target: "Bio", numerator: 1, denominator: 40},
		{source: "Wendigo", target: "Steel Orb", numerator: 1, denominator: 1},
		{source: "Steel Orb", target: "Demi", numerator: 1, denominator: 15},
		{source: "Torama", target: "Life Ring", numerator: 1, denominator: 1},
		{source: "Life Ring", target: "Life", numerator: 1, denominator: 20},
		{source: "Elastoid", target: "Steel Pipe", numerator: 1, denominator: 1},
		{source: "Steel Pipe", target: "Berserk", numerator: 1, denominator: 20},
		{source: "Grendel", target: "Dragon Fin", numerator: 1, denominator: 1},
		{source: "Dragon Fin", target: "Double", numerator: 1, denominator: 20},
		{source: "Blood Soul", target: "Zombie Powder", numerator: 1, denominator: 1},
		{source: "Zombie Powder", target: "Zombie", numerator: 1, denominator: 20},
		{source: "Imp", target: "Wizard Stone", numerator: 1, denominator: 1},
		{source: "Wizard Stone", target: "Stop", numerator: 1, denominator: 5},
		{source: "Spider Web", target: "Slow", numerator: 1, denominator: 20},
		{source: "Grat", target: "Magic Stone", numerator: 1, denominator: 1},
		{source: "Buel", target: "Magic Stone", numerator: 1, denominator: 1},
		{source: "Jelleye", target: "Magic Stone", numerator: 1, denominator: 1},
		{source: "Magic Stone", target: "Haste", numerator: 1, denominator: 5},
		{source: "Ochu", target: "Ochu Tentacle", numerator: 1, denominator: 1},
		{source: "Ochu Tentacle", target: "Blind", numerator: 1, denominator: 20},
		{source: "Chimera", target: "Regen Ring", numerator: 10, denominator: 1},
		{source: "Regen Ring", target: "Full-Life", numerator: 1, denominator: 20},
		{source: "Elnoyle", target: "Energy Crystal", numerator: 20, denominator: 2},
		{source: "Energy Crystal", target: "Pulse Ammo", numerator: 2, denominator: 20},
		{source: "Iron Giant", target: "Star Fragment", numerator: 9, denominator: 3},

		// Refinement Moment #3
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/the-timber-mission

		// Refinement Moment #4
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/galbadia-garden

		// Refinement Moment #5
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/deling-city

		// Refinement Moment #6
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/deling-city

		// Refinement Moment #7
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/the-escape

		// Refinement Moment #8
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/exploring-the-world

		// Refinement Moment #9
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/trabia-garden

		// Refinement Moment #10
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/the-aftermath

		// Refinement Moment #11
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/esthar

		// Refinement Moment #12
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/lunar-base

		// Refinement Moment #13
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/back-on-earth
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/the-final-mission

		// Refinement Moment #14
		// ref: https://gamefaqs.gamespot.com/ps4/266152-final-fantasy-viii-remastered/faqs/72431/final-preparations
	}
)
