package services

import (
	"backend/db/minitec_db"
	"backend/internal/http/services/project"
	"backend/internal/http/services/state"
	"backend/internal/http/services/station"
	"database/sql"
)

type Services struct {
	Project *project.Project
	Station *station.Station
	State   *state.State
}

func New(db *sql.DB, queries *minitec_db.Queries) *Services {
	station := station.New(db, queries)
	state := state.New(db, queries)
	return &Services{
		Project: project.New(db, queries, station, state),
		Station: station,
		State:   state,
	}
}
