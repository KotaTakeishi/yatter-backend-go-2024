package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	// Fetch statuses which has specified id
	FindByID(ctx context.Context, id int) (*object.Status, error)
	// Create a new status
	Create(ctx context.Context, tx *sqlx.Tx, status *object.Status) error
}
