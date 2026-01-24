package project

import (
	"backend/db/minitec_db"
	"backend/internal/http/csvparser"
	"backend/internal/http/services/state"
	"backend/internal/http/services/station"
	"context"
	"database/sql"
	"encoding/csv"
	"log/slog"
)

type Project struct {
	db      *sql.DB
	queries *minitec_db.Queries
	station *station.Station
	state   *state.State
}

func New(db *sql.DB, queries *minitec_db.Queries, station *station.Station, state *state.State) *Project {
	return &Project{
		db:      db,
		queries: queries,
		station: station,
		state:   state,
	}
}

func (p *Project) CreateProject(ctx context.Context, code string, name string) (*int64, error) {
	if name == "" {
		name = code
	}

	result, err := p.queries.CreateProject(ctx, minitec_db.CreateProjectParams{
		Code: code,
		Name: name,
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

func (p *Project) GetProject(ctx context.Context, id int64) (minitec_db.Project, error) {
	result, err := p.queries.GetProject(ctx, id)
	if err != nil {
		slog.Error("Failed to read from DB")
		return minitec_db.Project{}, err
	}

	return result, err
}

func (p *Project) UpdateProject() {

}

func (p *Project) DeleteProject(ctx context.Context, id int64) error {
	result, err := p.queries.DeleteProject(ctx, id)
	if err != nil {
		slog.Error("Failed to delete from DB")
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("Failed to get affected rows")
		return err
	}

	if rowsAffected == 0 {
		slog.Error("No rows affected")
		return sql.ErrNoRows
	}

	return nil
}

func (p *Project) ListProjects(ctx context.Context, limit int32, offset int32) ([]minitec_db.Project, error) {
	result, err := p.queries.ListProjects(ctx, minitec_db.ListProjectsParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		slog.Error("Failed to read from DB")
		return nil, err
	}

	return result, err
}

func (p *Project) ProjectHealth(ctx context.Context, id int64, reader *csv.Reader) (map[string][]int, error) {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	errs, err := csvparser.ParseCSV(ctx, reader, id, p.station, p.state, tx)
	if err != nil {
		return nil, err
	}

	if len(errs) > 0 {
		return errs, nil
	}

	_ = tx.Commit()
	return nil, nil
}
