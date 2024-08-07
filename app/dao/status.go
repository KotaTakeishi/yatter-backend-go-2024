package dao

import (
	"context"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Status
	status struct {
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

// Create status repository
func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

func (s *status) Create(ctx context.Context, tx *sqlx.Tx, status *object.Status) error {
	_, err := s.db.Exec("insert into status (account_id, content, create_at) values (?, ?, ?)",
		status.AccountID, status.Content, status.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert status: %w", err)
	}

	return nil
}
