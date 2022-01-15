package repositories

import (
	"fmt"

	"github.com/iamnande/cardmod/internal/database"
)

// rollback calls to tx.Rollback and wraps the given error
// with the rollback error if occurred.
func rollback(tx *database.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}
