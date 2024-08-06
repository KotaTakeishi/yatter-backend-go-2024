package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Account interface {
	Create(ctx context.Context, username, password string) (*CreateAccountDTO, error)
}

type account struct {
	db          *sqlx.DB
	accountRepo repository.Account
}

type CreateAccountDTO struct {
	Account *object.Account
}

type GetAccountDTO struct {
	Account *object.Account
}

var _ Account = (*account)(nil)

func NewAcocunt(db *sqlx.DB, accountRepo repository.Account) *account {
	return &account{
		db:          db,
		accountRepo: accountRepo,
	}
}

func (a *account) Create(ctx context.Context, username, password string) (*CreateAccountDTO, error) {
	acc, err := object.NewAccount(username, password)
	if err != nil {
		return nil, err
	}

	tx, err := a.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	if err := a.accountRepo.Create(ctx, tx, acc); err != nil {
		return nil, err
	}

	return &CreateAccountDTO{
		Account: acc,
	}, nil
}
