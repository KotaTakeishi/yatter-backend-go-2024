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

// FindByID : IDからユーザを取得
func (a *account) FindByID(ctx context.Context, id int64) (*object.Account, error) {
	entity := new(object.Account)
	err := a.db.QueryRowxContext(ctx, "select * from account where id = ?", id).StructScan(entity)
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

// Update : アカウント情報を更新
func (a *account) Update(ctx context.Context, tx *sqlx.Tx, acc *object.Account) error {
	_, err := tx.Exec("update account set display_name = ?, avatar = ?, header = ?, note = ? where id = ?",
		acc.DisplayName, acc.Avatar, acc.Header, acc.Note, acc.ID)
	if err != nil {
		return fmt.Errorf("failed to update account: %w", err)
	}

	return nil
}

// Follow : フォロー
func (a *account) Follow(ctx context.Context, tx *sqlx.Tx, followerID, followeeID int64) error {
	_, err := tx.Exec("insert into relationship (follower_id, followee_id) values (?, ?)", followerID, followeeID)
	if err != nil {
		return fmt.Errorf("failed to insert follow: %w", err)
	}

	return nil
}

func (a *account) GetRelationships(ctx context.Context, authUserID int64) ([]*object.Relationship, error) {
	entities := []*object.Relationship{}
	rows, err := a.db.QueryxContext(ctx, "select * from relationship where follower_id = ? or followee_id = ?", authUserID, authUserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to get relationships: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		entity := new(object.Relationship)
		err := rows.StructScan(entity)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		entities = append(entities, entity)
	}

	return entities, nil
}
