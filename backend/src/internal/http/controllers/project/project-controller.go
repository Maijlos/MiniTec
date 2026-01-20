package project

import (
	"backend/src/internal/http/controllers/project/dto/request"
	"backend/src/internal/http/controllers/project/dto/response"
	"backend/src/internal/http/messages"
	"backend/src/internal/http/services/project"
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Service *project.Project
}

func New(project *project.Project) *Project {
	return &Project{
		Service: project,
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
	id, err := p.Service.CreateProject(ctx, u.Code, u.Name)
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
	project, err := p.Service.GetProject(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, response.ErrorResponse{
				Code:         http.StatusNotFound,
				Message:      messages.ID_NOT_FOUND,
				ShortMessage: messages.ID_NOT_FOUND_SHORT,
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
	err = p.Service.DeleteProject(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, response.ErrorResponse{
				Code:         http.StatusNotFound,
				Message:      messages.ID_NOT_FOUND,
				ShortMessage: messages.ID_NOT_FOUND_SHORT,
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
	slog.Info("Listing projects")
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
	projects, err := p.Service.ListProjects(ctx, limit32, offset32)
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
