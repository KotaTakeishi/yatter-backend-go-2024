package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Timeline interface {
	FindPublicTimelines(ctx context.Context) (*GetTimelineDTO, error)
}

type timeline struct {
	db           *sqlx.DB
	timelineRepo repository.Timeline
}

type GetTimelineDTO struct {
	Timeline []*object.Status
}

var _ Timeline = (*timeline)(nil)

func NewTimeline(db *sqlx.DB, timelineRepo repository.Timeline) *timeline {
	return &timeline{
		db:           db,
		timelineRepo: timelineRepo,
	}
}

func (t *timeline) FindPublicTimelines(ctx context.Context) (*GetTimelineDTO, error) {
	timeline, err := t.timelineRepo.FindPublicTimelines(ctx)
	if err != nil {
		return nil, err
	}

	return &GetTimelineDTO{
		Timeline: timeline,
	}, nil
}
