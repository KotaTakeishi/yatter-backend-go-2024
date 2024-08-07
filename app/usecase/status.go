package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	FindByID(ctx context.Context, id int) (*GetStatusDTO, error)
	Create(ctx context.Context, account_id int, content string) (*CreateStatusDTO, error)
}

type status struct {
	db         *sqlx.DB
	statusRepo repository.Status
}

type CreateStatusDTO struct {
	Status struct {
		Status string `json:"status"`
	}
}

type GetStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status) *status {
	return &status{
		db:         db,
		statusRepo: statusRepo,
	}
}

func (s *status) FindByID(ctx context.Context, id int) (*GetStatusDTO, error) {
	status, err := s.statusRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetStatusDTO{
		Status: status,
	}, nil
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
		Status: struct {
			Status string `json:"status"`
		}{Status: status.Content},
	}, nil
}
