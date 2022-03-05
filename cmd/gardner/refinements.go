package main

import (
	"context"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/repositories"
)

// Refinement represents a refinement.
type Refinement struct {
	source      string
	target      string
	numerator   int32
	denominator int32
}

func (m *Refinement) Source() string {
	return m.source
}

func (m *Refinement) Target() string {
	return m.target
}

func (m *Refinement) Numerator() int32 {
	return m.numerator
}

func (m *Refinement) Denominator() int32 {
	return m.denominator
}

var (
	refinements = []*Refinement{
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
	}
)

func seedRefinements(ctx context.Context, db *database.Client) error {

	// initialize client
	refinementDB := repositories.NewRefinementRepository(db)

	// create the diff
	for i := 0; i < len(refinements); i++ {
		refinement, err := refinementDB.GetRefinement(ctx, refinements[i].Source(), refinements[i].Target())
		if err != nil {
			if !cerrors.IsAPIError(err) {
				return err
			}
		}

		if refinement == nil {
			_, err = refinementDB.CreateRefinement(ctx,
				refinements[i].Source(), refinements[i].Target(),
				refinements[i].Numerator(), refinements[i].Denominator(),
			)
			if err != nil {
				return err
			}
		}
	}

	// success
	return nil

}
