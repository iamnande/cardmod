package main

import (
	"context"

	"github.com/iamnande/cardmod/internal/cerrors"
	"github.com/iamnande/cardmod/internal/database"
	"github.com/iamnande/cardmod/internal/repositories"
)

// LimitBreak represents a limitbreak.
type LimitBreak struct {
	name string
}

func (c *LimitBreak) Name() string {
	return c.name
}

var (
	limitbreaks = []*LimitBreak{
		{name: "Ultra Waves"},
		{name: "Electrocute"},
		{name: "L?Death"},
		{name: "Degenerator"},
		{name: "Aqua Breath"},
		{name: "Micro Missiles"},
		{name: "Acid"},
		{name: "Gatling Gun"},
		{name: "Fire Breath"},
		{name: "Bad Breath"},
		{name: "White Wind"},
		{name: "Homing Laser"},
		{name: "Mighty Guard"},
		{name: "Ray-Bomb"},
		{name: "Shockwave Pulsar"},
	}
)

func seedLimitBreaks(ctx context.Context, db *database.Client) error {

	// initialize client
	limitbreakDB := repositories.NewLimitBreakRepository(db)

	// create the diff
	for i := 0; i < len(limitbreaks); i++ {
		limitbreak, err := limitbreakDB.GetLimitBreak(ctx, limitbreaks[i].Name())
		if err != nil {
			if !cerrors.IsAPIError(err) {
				return err
			}
		}

		if limitbreak == nil {
			_, err = limitbreakDB.CreateLimitBreak(ctx, limitbreaks[i].Name())
			if err != nil {
				return err
			}
		}
	}

	// success
	return nil

}
