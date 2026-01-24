package response

import "backend/db/minitec_db"

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
