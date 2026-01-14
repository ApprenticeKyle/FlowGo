package handler

import (
	"FLOWGO/internal/application/dto"
	"FLOWGO/internal/application/usecase"
	apperrors "FLOWGO/pkg/errors"

	"github.com/gin-gonic/gin"
)

type ProjectsHandler struct {
	BaseHandler
	createProjectUseCase *usecase.CreateProjectsCase
	updateProjectUseCase *usecase.UpdateProjectsCase
	getProjectUseCase    *usecase.GetProjectsCase
	listProjectsUseCase  *usecase.ListProjectsCase
	deleteProjectUseCase *usecase.DeleteProjectsCase
	projectsTeamUseCase  *usecase.ProjectsTeamUseCase
}

func NewProjectsHandler(
	createProjectUseCase *usecase.CreateProjectsCase,
	updateProjectUseCase *usecase.UpdateProjectsCase,
	getProjectUseCase *usecase.GetProjectsCase,
	listProjectsUseCase *usecase.ListProjectsCase,
	deleteProjectUseCase *usecase.DeleteProjectsCase,
	projectsTeamUseCase *usecase.ProjectsTeamUseCase,
) *ProjectsHandler {
	return &ProjectsHandler{
		createProjectUseCase: createProjectUseCase,
		updateProjectUseCase: updateProjectUseCase,
		getProjectUseCase:    getProjectUseCase,
		listProjectsUseCase:  listProjectsUseCase,
		deleteProjectUseCase: deleteProjectUseCase,
		projectsTeamUseCase:  projectsTeamUseCase,
	}
}

func (h *ProjectsHandler) CreateProject(c *gin.Context) {
	var req dto.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.HandleBadRequest(c, err.Error())
		return
	}

	project, err := h.createProjectUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		if appErr, ok := err.(*apperrors.AppError); ok {
			h.HandleError(c, appErr.Code, appErr.Message)
		} else {
			h.HandleInternalError(c, err.Error())
		}
		return
	}

	h.HandleSuccess(c, project)
}

func (h *ProjectsHandler) UpdateProject(c *gin.Context) {
	var req dto.UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.HandleBadRequest(c, err.Error())
		return
	}

	project, err := h.updateProjectUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		if appErr, ok := err.(*apperrors.AppError); ok {
			h.HandleError(c, appErr.Code, appErr.Message)
		} else {
			h.HandleInternalError(c, err.Error())
		}
		return
	}

	h.HandleSuccess(c, project)
}

func (h *ProjectsHandler) DeleteProject(c *gin.Context) {
	var req dto.DeleteProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.HandleBadRequest(c, err.Error())
		return
	}

	project, err := h.deleteProjectUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		if appErr, ok := err.(*apperrors.AppError); ok {
			h.HandleError(c, appErr.Code, appErr.Message)
		} else {
			h.HandleInternalError(c, err.Error())
		}
		return
	}

	h.HandleSuccess(c, project)
}

func (h *ProjectsHandler) GetProject(c *gin.Context) {
	var req dto.GetProjectRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.HandleBadRequest(c, err.Error())
		return
	}

	project, err := h.getProjectUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		if appErr, ok := err.(*apperrors.AppError); ok {
			h.HandleError(c, appErr.Code, appErr.Message)
		} else {
			h.HandleInternalError(c, err.Error())
		}
		return
	}

	h.HandleSuccess(c, project)
}

func (h *ProjectsHandler) ListProjects(c *gin.Context) {
	var req dto.PageRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.HandleBadRequest(c, err.Error())
		return
	}

	projects, err := h.listProjectsUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		if appErr, ok := err.(*apperrors.AppError); ok {
			h.HandleError(c, appErr.Code, appErr.Message)
		} else {
			h.HandleInternalError(c, err.Error())
		}
		return
	}

	h.HandleSuccessWithPage(c, projects.List, projects.Page.Page, projects.Page.PageSize, projects.Page.Total)
}

func (h *ProjectsHandler) ProjectTeams(c *gin.Context) {
	teams, err := h.projectsTeamUseCase.Execute(c.Request.Context())
	if err != nil {
		if appErr, ok := err.(*apperrors.AppError); ok {
			h.HandleError(c, appErr.Code, appErr.Message)
		} else {
			h.HandleInternalError(c, err.Error())
		}
		return
	}

	h.HandleSuccess(c, teams)
}
