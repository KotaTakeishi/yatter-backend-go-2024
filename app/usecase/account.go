package usecase

import (
	"context"
	"log"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Account interface {
	FindByUsername(ctx context.Context, username string) (*GetAccountDTO, error)
	Create(ctx context.Context, username, password string) (*CreateAccountDTO, error)
	Update(ctx context.Context, id int64, display_name, note, avatar, header *string) (*UpdateAccountDTO, error)
	Follow(ctx context.Context, followerID, followeeID int64) (*FollowAccountDTO, error)
}

type account struct {
	db          *sqlx.DB
	accountRepo repository.Account
}

type GetAccountDTO struct {
	Account *object.Account
}

type CreateAccountDTO struct {
	Account *object.Account
}

type UpdateAccountDTO struct {
	Account *object.Account
}

type FollowAccountDTO struct {
	Relation *struct {
		ID           int64 `json:"id"`
		Following    bool  `json:"following"`
		Followerd_by bool  `json:"followerd_by"`
	}
}

var _ Account = (*account)(nil)

func NewAcocunt(db *sqlx.DB, accountRepo repository.Account) *account {
	return &account{
		db:          db,
		accountRepo: accountRepo,
	}
}

func (a *account) FindByUsername(ctx context.Context, username string) (*GetAccountDTO, error) {
	acc, err := a.accountRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return &GetAccountDTO{
		Account: acc,
	}, nil
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
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("tx.Rollback() failed: %v", rbErr)
			}
		}

		if err := tx.Commit(); err != nil {
			log.Printf("tx.Commit() failed: %v", err)
		}
	}()

	if err := a.accountRepo.Create(ctx, tx, acc); err != nil {
		return nil, err
	}

	return &CreateAccountDTO{
		Account: acc,
	}, nil
}

func (a *account) Update(ctx context.Context, id int64, display_name, note, avatar, header *string) (*UpdateAccountDTO, error) {
	acc, err := a.accountRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if display_name != nil {
		acc.DisplayName = display_name
	}
	if note != nil {
		acc.Note = note
	}
	if avatar != nil {
		acc.Avatar = avatar
	}
	if header != nil {
		acc.Header = header
	}

	tx, err := a.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				log.Printf("tx.Rollback() failed: %v", rbErr)
			}
		}

		if err := tx.Commit(); err != nil {
			log.Printf("tx.Commit() failed: %v", err)
		}
	}()

	if err := a.accountRepo.Update(ctx, tx, acc); err != nil {
		return nil, err
	}

	return &UpdateAccountDTO{
		Account: acc,
	}, nil
}

func (a *account) Follow(ctx context.Context, followerID, followeeID int64) (*FollowAccountDTO, error) {
	tx, err := a.db.Beginx()
	if err != nil {
		return nil, err
	}

	if err := a.accountRepo.Follow(ctx, tx, followerID, followeeID); err != nil {
		return nil, err
	}

	if err := recover(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			log.Printf("tx.Rollback() failed: %v", rbErr)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Printf("tx.Commit() failed: %v", err)
	}

	relationships, err := a.accountRepo.GetRelationships(ctx, followerID)
	if err != nil {
		return nil, err
	}

	following := false
	followerd_by := false
	for _, r := range relationships {
		if r.FollowerID == followerID && r.FolloweeID == followeeID {
			following = true
		}
		if r.FollowerID == followeeID && r.FolloweeID == followerID {
			followerd_by = true
		}
	}

	return &FollowAccountDTO{
		Relation: &struct {
			ID           int64 `json:"id"`
			Following    bool  `json:"following"`
			Followerd_by bool  `json:"followerd_by"`
		}{
			ID:           followeeID,
			Following:    following,
			Followerd_by: followerd_by,
		},
	}, nil
}
