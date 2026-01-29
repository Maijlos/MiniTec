package response

import (
	"backend/db/minitec_db"
	"backend/internal/http/services/project"
)


type Project struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type SuccessfulResponse struct {
	Code         int       `json:"code"`
	ShortMessage string    `json:"short_message"`
	Data         []Project `json:"data"`
}

type SuccessfulResponseHealth struct {
	Code         int       `json:"code"`
	ShortMessage string    `json:"short_message"`
	Data         map[string][]project.StringState `json:"data"`
}

type ErrorResponse struct {
	Code         int    `json:"code"`
	ShortMessage string `json:"short_message"`
	Message      string `json:"message"`
}

type ErrorParsingCSV struct {
	Code     int              `json:"code"`
	Messages string           `json:"messages"`
	Errors   map[string][]int `json:"errors"`
}

func MapModelsToResponse(ps []minitec_db.Project) []Project {
	var projects []Project
	for _, p := range ps {
		projects = append(projects, MapModelToResponse(p))
	}
	return projects
}

func MapModelToResponse(p minitec_db.Project) Project {
	return Project{
		Id:   p.ID,
		Code: p.Code,
		Name: p.Name,
	}
}
