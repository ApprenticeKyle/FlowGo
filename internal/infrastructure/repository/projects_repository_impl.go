package repository

import (
	"FLOWGO/internal/domain/entity"
	"FLOWGO/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

// projectsRepositoryImpl 项目仓储实现
type projectsRepositoryImpl struct {
	db *gorm.DB
}

// NewProjectsRepository 创建项目仓储实例
func NewProjectsRepository(db *gorm.DB) repository.ProjectsRepository {
	return &projectsRepositoryImpl{
		db: db,
	}
}

// Create 创建项目
func (r *projectsRepositoryImpl) Create(ctx context.Context, project *entity.Project) error {
	return r.db.WithContext(ctx).Create(project).Error
}

// Delete 删除项目
func (r *projectsRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	return r.db.WithContext(ctx).Delete(&entity.Project{}, id).Error
}

// FindById 根据ID查找
func (r *projectsRepositoryImpl) FindByID(ctx context.Context, id uint64) (*entity.Project, error) {
	var project entity.Project
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&project).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// List 列表查询
func (r *projectsRepositoryImpl) List(ctx context.Context, page, pageSize int) ([]*entity.Project, int64, error) {
	var projects []*entity.Project
	var total int64

	offset := (page - 1) * pageSize

	// 查询总数
	if err := r.db.WithContext(ctx).Model(&entity.Project{}).
		Where("deleted_at IS NULL").
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询列表
	if err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&projects).Error; err != nil {
		return nil, 0, err
	}

	return projects, total, nil
}

// Update 更新项目
func (r *projectsRepositoryImpl) Update(ctx context.Context, project *entity.Project) error {
	return r.db.WithContext(ctx).Save(project).Error
}

// ListAvailableTeams 列表查询可用团队
func (r *projectsRepositoryImpl) ListAvailableTeams(ctx context.Context) ([]*entity.Team, error) {
	var teams []*entity.Team
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}
