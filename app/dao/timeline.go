package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Timeline
	timeline struct {
		db *sqlx.DB
	}
)

var _ repository.Timeline = (*timeline)(nil)

// Create timeline repository
func NewTimeline(db *sqlx.DB) *timeline {
	return &timeline{db: db}
}

func (t *timeline) FindPublicTimelines(ctx context.Context) ([]*object.Status, error) {
	entities := []*object.Status{}
	rows, err := t.db.QueryxContext(ctx, "select * from status")

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to get public timelines: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var entity object.Status
		err := rows.StructScan(&entity)
		if err != nil {
			log.Printf("failed to scan row: %v", err)
			continue
		}
		entities = append(entities, &entity)
	}

	return entities, nil
}
