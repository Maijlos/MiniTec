package station

import (
	"backend/db/minitec_db"
	"context"
	"database/sql"
	"log/slog"
)

type Station struct {
	Db      *sql.DB
	Queries *minitec_db.Queries
}

func New(db *sql.DB, queries *minitec_db.Queries) *Station {
	return &Station{
		Db:      db,
		Queries: queries,
	}
}

func (s *Station) CreateStation(ctx context.Context, projectId int64, name string, tx *sql.Tx) (*int64, error) {
	qtx := s.Queries.WithTx(tx)
	result, err := qtx.CreateStation(ctx, minitec_db.CreateStationParams{
		ProjectID: projectId,
		Name:      name,
	})

	if err != nil {
		slog.Error("Failed to write to DB")
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		slog.Error("Failed to get id")
		return nil, err
	}

	return &id, nil
}

func (s *Station) GetStationId(ctx context.Context, projectId int64, name string, tx *sql.Tx) (*int64, error) {
	qtx := s.Queries.WithTx(tx)
	id, err := qtx.GetStationId(ctx, minitec_db.GetStationIdParams{
		ProjectID: projectId,
		Name:      name,
	})
	if err != nil {
		slog.Error("Failed to read from DB")
		return nil, err
	}

	return &id, err
}
