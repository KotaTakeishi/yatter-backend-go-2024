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
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}
)

var _ repository.Account = (*account)(nil)

// Create accout repository
func NewAccount(db *sqlx.DB) *account {
	return &account{db: db}
}

// FindByUsername : ユーザ名からユーザを取得
func (a *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := a.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

func (a *account) Create(ctx context.Context, tx *sqlx.Tx, acc *object.Account) error {
	_, err := a.db.Exec("insert into account (username, password_hash, display_name, avatar, header, note, create_at) values (?, ?, ?, ?, ?, ?, ?)",
		acc.Username, acc.PasswordHash, acc.DisplayName, acc.Avatar, acc.Header, acc.Note, acc.CreateAt)
	if err != nil {
		return fmt.Errorf("failed to insert account: %w", err)
	}

	return nil
}
