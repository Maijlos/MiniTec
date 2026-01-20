package project

import (
	"backend/db/minitec_db"
	"context"
	"database/sql"
	"log/slog"
)

type Project struct {
	queries *minitec_db.Queries
}

func New(queries *minitec_db.Queries) *Project {
	return &Project{
		queries: queries,
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
