package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	// Fetch account which has specified ID
	FindByID(ctx context.Context, id int64) (*object.Account, error)
	// Create a new account
	Create(ctx context.Context, tx *sqlx.Tx, acc *object.Account) error
	// Update account
	Update(ctx context.Context, tx *sqlx.Tx, acc *object.Account) error
	// Follow account
	Follow(ctx context.Context, tx *sqlx.Tx, followerID, followeeID int64) error
	// Fetch relationships
	GetRelationships(ctx context.Context, authUserID int64) ([]*object.Relationship, error)
}
