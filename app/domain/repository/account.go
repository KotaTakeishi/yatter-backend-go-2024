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
	// エンドポイントでは使用しないが，IDからstatusを取得する際に必要
	FindByID(ctx context.Context, id int64) (*object.Account, error)
	// Create a new account
	Create(ctx context.Context, tx *sqlx.Tx, acc *object.Account) error
}
