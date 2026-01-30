package project

import (
	"backend/internal/http/controllers/project/dto/request"
	"backend/internal/http/controllers/project/dto/response"
	"backend/internal/http/messages"
	projectService "backend/internal/http/services/project"
	"database/sql"
	"encoding/csv"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Project struct {
	ProjectService *projectService.Project
}

func New(ProjectService *projectService.Project) *Project {
	return &Project{
		ProjectService: ProjectService,
	}
}

func (p *Project) CreateProject(c echo.Context) error {
	u := new(request.CreateProject)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:         http.StatusBadRequest,
			Message:      messages.INVALID_JSON,
			ShortMessage: messages.INVALID_JSON_SHORT,
		})
	}

	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:         http.StatusBadRequest,
			Message:      messages.INVALID_BODY,
			ShortMessage: messages.INVALID_BODY_SHORT,
		})
	}

	ctx := c.Request().Context()
	id, err := p.ProjectService.CreateProject(ctx, u.Code, u.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:         http.StatusInternalServerError,
			Message:      messages.SERVER_ERROR,
			ShortMessage: messages.SERVER_ERROR_SHORT,
		})
	}

	return c.JSON(http.StatusCreated, response.SuccessfulResponse{
		Code:         http.StatusCreated,
		ShortMessage: messages.SUCCESS,
		Data: []response.Project{
			{
				Id:   *id,
				Code: u.Code,
				Name: u.Name,
			},
		},
	})
}

func (p *Project) GetProject(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:         http.StatusBadRequest,
			Message:      messages.INVALID_ID,
			ShortMessage: messages.INVALID_ID_SHORT,
		})
	}

	ctx := c.Request().Context()
	project, err := p.ProjectService.GetProject(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, response.ErrorResponse{
				Code:         http.StatusNotFound,
				Message:      messages.NOT_FOUND,
				ShortMessage: messages.NOT_FOUND_SHORT,
			})
		}

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:         http.StatusInternalServerError,
			Message:      messages.SERVER_ERROR,
			ShortMessage: messages.SERVER_ERROR_SHORT,
		})
	}

	return c.JSON(http.StatusOK, response.SuccessfulResponse{
		Code:         http.StatusOK,
		ShortMessage: messages.SUCCESS,
		Data:         []response.Project{response.MapModelToResponse(project)},
	})
}

func (p *Project) GetProjectByCode(c echo.Context) error {
	u := new(request.GetProjectByCode)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:         http.StatusBadRequest,
			Message:      messages.INVALID_JSON,
			ShortMessage: messages.INVALID_JSON_SHORT,
		})
	}

	if err := c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:         http.StatusBadRequest,
			Message:      messages.INVALID_BODY,
			ShortMessage: messages.INVALID_BODY_SHORT,
		})
	}

	ctx := c.Request().Context()
	project, err := p.ProjectService.GetProjectByCode(ctx, u.Code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, response.ErrorResponse{
				Code:         http.StatusNotFound,
				Message:      messages.NOT_FOUND,
				ShortMessage: messages.NOT_FOUND_SHORT,
			})
		}

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:         http.StatusInternalServerError,
			Message:      messages.SERVER_ERROR,
			ShortMessage: messages.SERVER_ERROR_SHORT,
		})
	}

	return c.JSON(http.StatusOK, response.SuccessfulResponse{
		Code: http.StatusOK,
		ShortMessage: messages.SUCCESS,
		Data: []response.Project{response.MapModelToResponse(project)},
	})
}

func (p *Project) UpdateProject() {

}

func (p *Project) DeleteProject(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:         http.StatusBadRequest,
			Message:      messages.INVALID_ID,
			ShortMessage: messages.INVALID_ID_SHORT,
		})
	}

	ctx := c.Request().Context()
	err = p.ProjectService.DeleteProject(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, response.ErrorResponse{
				Code:         http.StatusNotFound,
				Message:      messages.NOT_FOUND,
				ShortMessage: messages.NOT_FOUND_SHORT,
			})
		}

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:         http.StatusInternalServerError,
			Message:      messages.SERVER_ERROR,
			ShortMessage: messages.SERVER_ERROR_SHORT,
		})
	}

	return c.JSON(http.StatusOK, response.SuccessfulResponse{
		Code:         http.StatusOK,
		ShortMessage: messages.SUCCESS,
		Data:         []response.Project{},
	})
}

func (p *Project) ListProjects(c echo.Context) error {
	limit := int64(10)
	offset := int64(0)
	limitString := c.QueryParam("limit")
	if limitString != "" {
		var err error
		limit, err = strconv.ParseInt(limitString, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Code:         http.StatusBadRequest,
				Message:      messages.INVALID_LIMIT,
				ShortMessage: messages.INVALID_LIMIT_SHORT,
			})
		}
	}
	offsetString := c.QueryParam("offset")
	if offsetString != "" {
		var err error
		offset, err = strconv.ParseInt(offsetString, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				Code:         http.StatusBadRequest,
				Message:      messages.INVALID_OFFSET,
				ShortMessage: messages.INVALID_OFFSET_SHORT,
			})
		}
	}

	ctx := c.Request().Context()

	limit32 := int32(limit)
	offset32 := int32(offset)
	projects, err := p.ProjectService.ListProjects(ctx, limit32, offset32)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:         http.StatusInternalServerError,
			Message:      messages.SERVER_ERROR,
			ShortMessage: messages.SERVER_ERROR_SHORT,
		})
	}

	return c.JSON(http.StatusOK, response.SuccessfulResponse{
		Code:         http.StatusOK,
		ShortMessage: messages.SUCCESS,
		Data:         response.MapModelsToResponse(projects),
	})
}

func (p *Project) ProjectHealth(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:         http.StatusBadRequest,
			Message:      messages.INVALID_ID,
			ShortMessage: messages.INVALID_ID_SHORT,
		})
	}

	ctx := c.Request().Context()
	project, err := p.ProjectService.GetProject(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, response.ErrorResponse{
				Code:         http.StatusNotFound,
				Message:      messages.NOT_FOUND,
				ShortMessage: messages.NOT_FOUND_SHORT,
			})
		}

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:         http.StatusInternalServerError,
			Message:      messages.SERVER_ERROR,
			ShortMessage: messages.SERVER_ERROR_SHORT,
		})
	}

	body := c.Request().Body
	defer func() {
		_ = body.Close()
	}()

	reader := csv.NewReader(body)
	errs, err := p.ProjectService.ProjectHealth(ctx, project.ID, reader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:         http.StatusInternalServerError,
			Message:      messages.SERVER_ERROR,
			ShortMessage: messages.SERVER_ERROR_SHORT,
		})
	}

	if errs != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorParsingCSV{
			Code:     http.StatusBadRequest,
			Messages: "Invalid CSV file",
			Errors:   errs,
		})
	}

	return c.JSON(http.StatusOK, response.SuccessfulResponse{
		Code:         http.StatusOK,
		ShortMessage: messages.SUCCESS,
		Data:         []response.Project{},
	})
}

func (p *Project) GetProjectHealth(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:         http.StatusBadRequest,
			Message:      messages.INVALID_ID,
			ShortMessage: messages.INVALID_ID_SHORT,
		})
	}

	ctx := c.Request().Context()
	project, err := p.ProjectService.GetProject(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, response.ErrorResponse{
				Code:         http.StatusNotFound,
				Message:      messages.NOT_FOUND,
				ShortMessage: messages.NOT_FOUND_SHORT,
			})
		}

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:         http.StatusInternalServerError,
			Message:      messages.SERVER_ERROR,
			ShortMessage: messages.SERVER_ERROR_SHORT,
		})
	}

	result, err :=p.ProjectService.GetProjectHealth(ctx, project.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:         http.StatusInternalServerError,
			Message:      messages.SERVER_ERROR,
			ShortMessage: messages.SERVER_ERROR_SHORT,
		})
	}

	return c.JSON(http.StatusOK, response.SuccessfulResponseHealth{
		Code:         http.StatusOK,
		ShortMessage: messages.SUCCESS,
		Data:         result,
	})
}
