package request

type CreateProject struct {
	Code string `json:"code" validate:"required"`
	Name string `json:"name"`
}
