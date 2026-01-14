package repository

import (
	"FLOWGO/internal/domain/entity"
	"FLOWGO/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

type projectsTeamsRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectsTeamsRepository(db *gorm.DB) repository.ProjectsTeamsRepository {
	return &projectsTeamsRepositoryImpl{
		db: db,
	}
}

// DeleteByProjectId 删除项目团队
func (r *projectsTeamsRepositoryImpl) DeleteByProjectId(ctx context.Context, projectId uint64) error {
	return r.db.WithContext(ctx).Where("project_id = ?", projectId).Delete(&entity.ProjectsTeams{}).Error
}

// SaveBatch 批量保存项目团队
func (r *projectsTeamsRepositoryImpl) SaveBatch(ctx context.Context, entities []*entity.ProjectsTeams) error {
	return r.db.WithContext(ctx).Create(entities).Error
}
