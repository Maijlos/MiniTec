package services

import (
	"backend/db/minitec_db"
	project "backend/src/internal/http/services/project"
)

type Services struct {
	Project *project.Project
}

func New(queries *minitec_db.Queries) *Services {
	return &Services{
		Project: project.New(queries),
	}
}
