package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	Create(ctx context.Context, account_id int, content string) (*CreateStatusDTO, error)
}

type status struct {
	db         *sqlx.DB
	statusRepo repository.Status
}

type CreateStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:         db,
		statusRepo: statusRepo,
	}
}

func (s *status) Create(ctx context.Context, account_id int, content string) (*CreateStatusDTO, error) {
	status := object.NewStatus(account_id, content)

	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}

		tx.Commit()
	}()

	if err := s.statusRepo.Create(ctx, tx, status); err != nil {
		return nil, err
	}

	return &CreateStatusDTO{
		Status: status,
	}, nil
}
