package refinements

var (
	// refinement is the collection of refinement entities.
	refinements = [26]*refinement{
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

		// Initial Magic
		{source: "Abyss Worm", target: "Windmill", numerator: 1, denominator: 1},
		{source: "Windmill", target: "Tornado", numerator: 1, denominator: 20},
		{source: "Ruby Dragon", target: "Inferno Fang", numerator: 10, denominator: 1},
		{source: "Inferno Fang", target: "Flare", numerator: 1, denominator: 20},
		{source: "Hexadrago", target: "Red Fang", numerator: 3, denominator: 1},
		{source: "Red Fang", target: "Firaga", numerator: 1, denominator: 20},
	}
)
