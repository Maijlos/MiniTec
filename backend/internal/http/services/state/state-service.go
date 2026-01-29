package state

import (
	"backend/db/minitec_db"
	"context"
	"database/sql"
	"log/slog"
	"time"
)

type State struct {
	db      *sql.DB
	queries *minitec_db.Queries
}

func New(db *sql.DB, queries *minitec_db.Queries) *State {
	return &State{
		db:      db,
		queries: queries,
	}
}

func (s *State) CreateState(ctx context.Context, tx *sql.Tx, stationId int64, startDate time.Time, endDate time.Time, finalStatus int32) (*int64, error) {
	qtx := s.queries.WithTx(tx)
	result, err := qtx.CreateState(ctx, minitec_db.CreateStateParams{
		FinalState: finalStatus,
		StartDate:  sql.NullTime{Time: startDate, Valid: !startDate.IsZero()},
		EndDate:    sql.NullTime{Time: endDate, Valid: !endDate.IsZero()},
		StationID:  stationId,
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

func (s *State) GetAllStatesByStation(ctx context.Context, stationId int64) ([]minitec_db.State, error) {
	result, err := s.queries.ListStationsByStation(ctx, stationId)
	if err != nil {
		slog.Error("Failed to read from DB")
		return nil, err
	}
	return result, nil
}
