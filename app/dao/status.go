package dao

import (
	"context"
	"database/sql"
	"errors"
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

func (s *status) FindByID(ctx context.Context, id int64) (*object.Status, error) {
	entity := new(object.Status)
	err := s.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}

	return entity, nil
}

func (s *status) Create(ctx context.Context, tx *sqlx.Tx, status *object.Status) error {
	_, err := s.db.Exec("insert into status (account_id, content, create_at) values (?, ?, ?)",
		status.AccountID, status.Content, status.CreateAt)
	if err != nil {
		return fmt.Errorf("failed to insert status: %w", err)
	}

	return nil
}
