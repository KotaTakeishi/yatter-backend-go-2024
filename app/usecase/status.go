package usecase

import (
	"context"
	"time"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	FindByID(ctx context.Context, id int64) (*GetStatusDTO, error)
	Create(ctx context.Context, account_id int64, content string) (*CreateStatusDTO, error)
}

type status struct {
	db           *sqlx.DB
	acccountRepo repository.Account
	statusRepo   repository.Status
}

type CreateStatusDTO struct {
	Status *struct {
		Status string `json:"status"`
	}
}

type GetStatusDTO struct {
	Status *struct {
		ID       int64           `json:"id"`
		Account  *object.Account `json:"account"`
		Content  string          `json:"content"`
		CreateAt time.Time       `json:"create_at"`
	}
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, accountRepo repository.Account, statusRepo repository.Status) *status {
	return &status{
		db:           db,
		acccountRepo: accountRepo,
		statusRepo:   statusRepo,
	}
}

func (s *status) FindByID(ctx context.Context, id int64) (*GetStatusDTO, error) {
	status, err := s.statusRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if status == nil {
		return nil, nil
	}

	account, err := s.acccountRepo.FindByID(ctx, status.AccountID)
	if err != nil {
		return nil, err
	}

	return &GetStatusDTO{
		Status: &struct {
			ID       int64           `json:"id"`
			Account  *object.Account `json:"account"`
			Content  string          `json:"content"`
			CreateAt time.Time       `json:"create_at"`
		}{
			ID:       status.ID,
			Account:  account,
			Content:  status.Content,
			CreateAt: status.CreateAt,
		},
	}, nil
}

func (s *status) Create(ctx context.Context, account_id int64, content string) (*CreateStatusDTO, error) {
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
		Status: &struct {
			Status string `json:"status"`
		}{Status: status.Content},
	}, nil
}
