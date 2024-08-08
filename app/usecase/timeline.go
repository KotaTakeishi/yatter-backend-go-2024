package usecase

import (
	"context"
	"time"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Timeline interface {
	FindPublicTimelines(ctx context.Context) (*GetTimelineDTOs, error)
}

type timeline struct {
	db           *sqlx.DB
	accountRepo  repository.Account
	timelineRepo repository.Timeline
}

type GetTimelineDTO struct {
	ID       int64           `json:"id"`
	Account  *object.Account `json:"account"`
	Content  string          `json:"content"`
	CreateAt time.Time       `json:"create_at"`
}

type GetTimelineDTOs struct {
	Timeline []*GetTimelineDTO
}

var _ Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB, accountRepo repository.Account, timelineRepo repository.Timeline) *timeline {
	return &timeline{
		db:           db,
		accountRepo:  accountRepo,
		timelineRepo: timelineRepo,
	}
}

func (t *timeline) FindPublicTimelines(ctx context.Context) (*GetTimelineDTOs, error) {
	timeline, err := t.timelineRepo.FindPublicTimelines(ctx)
	if err != nil {
		return nil, err
	}

	dtos := []*GetTimelineDTO{}
	for _, v := range timeline {
		// TODO: for文でクエリ叩いているのはいかがなものか
		// Layered Architectureに従った結果ではある
		account, err := t.accountRepo.FindByID(ctx, v.AccountID)
		if err != nil {
			return nil, err
		}
		dto := GetTimelineDTO{
			ID:       v.ID,
			Account:  account,
			Content:  v.Content,
			CreateAt: v.CreateAt,
		}
		dtos = append(dtos, &dto)
	}

	return &GetTimelineDTOs{Timeline: dtos}, nil
}
