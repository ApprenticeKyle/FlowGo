package usecase

import (
	"context"
	"errors"

	"FLOWGO/internal/application/dto"
	"FLOWGO/internal/domain/entity"
	"FLOWGO/internal/domain/repository"
	"FLOWGO/pkg/utils"
)

type CreateProjectsCase struct {
	projectRepo repository.ProjectsRepository
}

func NewCreateProjectsCase(projectRepo repository.ProjectsRepository) *CreateProjectsCase {
	return &CreateProjectsCase{
		projectRepo: projectRepo,
	}
}

func (uc *CreateProjectsCase) Execute(ctx context.Context, req dto.CreateProjectRequest) (*dto.CreateProjectResponse, error) {
	project := &entity.Project{
		Name:        req.Name,
		Description: req.Description,
		OwnerId:     req.OwnerId,
	}
	err := uc.projectRepo.Create(ctx, project)
	if err != nil {
		return nil, errors.New("创建项目失败")
	}
	return &dto.CreateProjectResponse{
		ID: project.ID,
	}, nil
}

type UpdateProjectsCase struct {
	projectRepo       repository.ProjectsRepository
	projectsTeamsRepo repository.ProjectsTeamsRepository
}

func NewUpdateProjectsCase(projectRepo repository.ProjectsRepository) *UpdateProjectsCase {
	return &UpdateProjectsCase{
		projectRepo: projectRepo,
	}
}

func (uc *UpdateProjectsCase) Execute(ctx context.Context, req dto.UpdateProjectRequest) (*dto.UpdateProjectResponse, error) {
	// 先查找现有项目
	existingProject, err := uc.projectRepo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, errors.New("查找项目失败")
	}
	if existingProject == nil {
		return nil, errors.New("项目不存在")
	}

	// 更新项目信息
	project := &entity.Project{
		BaseEntity: entity.BaseEntity{
			ID: req.ID,
		},
		OwnerId:     req.OwnerId,
		Status:      req.Status,
		Deadline:    req.Deadline.Time,  // 从 utils.Time 转换为 time.Time
		StartDate:   req.StartDate.Time, // 从 utils.Time 转换为 time.Time
		Members:     req.Members,
		Progress:    req.Progress,
		Priority:    req.Priority,
		Starred:     req.Starred,
		Archived:    req.Archived,
		CoverImage:  req.CoverImage,
		Name:        req.Name,
		Description: req.Description,
	}
	err = uc.projectRepo.Update(ctx, project)
	if err != nil {
		return nil, errors.New("更新项目失败")
	}
	//先删除
	err = uc.projectsTeamsRepo.DeleteByProjectId(ctx, req.ID)
	if err != nil {
		return nil, errors.New("删除项目团队失败")
	}
	//再保存
	entities := make([]*entity.ProjectsTeams, 0, len(req.TeamIds))
	for _, teamId := range req.TeamIds {
		entities = append(entities, &entity.ProjectsTeams{
			ProjectId: project.ID,
			TeamId:    teamId,
		})
	}
	err = uc.projectsTeamsRepo.SaveBatch(ctx, entities)
	if err != nil {
		return nil, errors.New("保存项目团队失败")
	}
	return &dto.UpdateProjectResponse{
		ID:          project.ID,
		TeamIds:     req.TeamIds,
		Tags:        req.Tags,
		Members:     req.Members,
		Progress:    req.Progress,
		Priority:    req.Priority,
		Starred:     req.Starred,
		Archived:    req.Archived,
		CoverImage:  req.CoverImage,
		Deadline:    req.Deadline,
		StartDate:   req.StartDate,
		Name:        req.Name,
		Description: req.Description,
		OwnerId:     req.OwnerId,
		Status:      req.Status,
	}, nil
}

type DeleteProjectsCase struct {
	projectRepo repository.ProjectsRepository
}

func NewDeleteProjectsCase(projectRepo repository.ProjectsRepository) *DeleteProjectsCase {
	return &DeleteProjectsCase{
		projectRepo: projectRepo,
	}
}

func (uc *DeleteProjectsCase) Execute(ctx context.Context, req dto.DeleteProjectRequest) (*dto.DeleteProjectResponse, error) {
	project, err := uc.projectRepo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, errors.New("删除项目失败")
	}
	if project == nil {
		return nil, errors.New("项目不存在")
	}
	err = uc.projectRepo.Delete(ctx, req.ID)
	if err != nil {
		return nil, errors.New("删除项目失败")
	}
	return &dto.DeleteProjectResponse{
		ID: project.ID,
	}, nil
}

type GetProjectsCase struct {
	projectRepo repository.ProjectsRepository
}

func NewGetProjectsCase(projectRepo repository.ProjectsRepository) *GetProjectsCase {
	return &GetProjectsCase{
		projectRepo: projectRepo,
	}
}

func (uc *GetProjectsCase) Execute(ctx context.Context, req dto.GetProjectRequest) (*dto.GetProjectResponse, error) {
	project, err := uc.projectRepo.FindByID(ctx, req.ID)
	if err != nil {
		return nil, errors.New("获取项目失败")
	}
	if project == nil {
		return nil, errors.New("项目不存在")
	}
	return &dto.GetProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		OwnerId:     project.OwnerId,
		Status:      project.Status,
		Deadline:    utils.NewTime(project.Deadline),
		StartDate:   utils.NewTime(project.StartDate),
		Members:     project.Members,
		Progress:    project.Progress,
		Priority:    project.Priority,
		Starred:     project.Starred,
		Archived:    project.Archived,
		CoverImage:  project.CoverImage,
	}, nil
}

type ListProjectsCase struct {
	projectRepo repository.ProjectsRepository
}

func NewListProjectsCase(projectRepo repository.ProjectsRepository) *ListProjectsCase {
	return &ListProjectsCase{
		projectRepo: projectRepo,
	}
}

func (uc *ListProjectsCase) Execute(ctx context.Context, req dto.PageRequest) (*dto.ProjectListResponse, error) {
	projects, total, err := uc.projectRepo.List(ctx, req.Page, req.GetPageSize())
	if err != nil {
		return nil, errors.New("获取项目列表失败")
	}
	projectResponses := make([]*dto.ProjectResponse, 0, len(projects))
	for _, project := range projects {
		projectResponses = append(projectResponses, &dto.ProjectResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			OwnerId:     project.OwnerId,
			Status:      project.Status,
			Deadline:    utils.NewTime(project.Deadline),
			StartDate:   utils.NewTime(project.StartDate),
			Members:     project.Members,
			Progress:    project.Progress,
			Priority:    project.Priority,
		})
	}
	return &dto.ProjectListResponse{
		List: projectResponses,
		Page: dto.PageResponse{
			Page:     req.Page,
			PageSize: req.GetPageSize(),
			Total:    total,
		},
	}, nil
}

type ProjectsTeamUseCase struct {
	teamRepo repository.TeamRepository
}

func NewProjectsTeamUseCase(teamRepo repository.TeamRepository) *ProjectsTeamUseCase {
	return &ProjectsTeamUseCase{
		teamRepo: teamRepo,
	}
}

func (uc *ProjectsTeamUseCase) Execute(ctx context.Context) (*dto.ProjectTeamsResponse, error) {
	teams, err := uc.teamRepo.ListAvailableTeams(ctx)
	if err != nil {
		return nil, errors.New("获取项目团队失败")
	}
	teamResponses := make([]*dto.TeamResponse, 0, len(teams))
	for _, team := range teams {
		teamResponses = append(teamResponses, &dto.TeamResponse{
			ID:   team.ID,
			Name: team.Name,
		})
	}
	return &dto.ProjectTeamsResponse{
		Teams: teamResponses,
	}, nil
}
