package repository

import (
	"FLOWGO/internal/domain/entity"
	"FLOWGO/internal/domain/repository"
	"context"

	"gorm.io/gorm"
)

type teamRepositoryImpl struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) repository.TeamRepository {
	return &teamRepositoryImpl{
		db: db,
	}
}

// ListAvailableTeams 列表查询可用团队
func (r *teamRepositoryImpl) ListAvailableTeams(ctx context.Context) ([]*entity.Team, error) {
	var teams []*entity.Team
	err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Find(&teams).Error
	if err != nil {
		return nil, err
	}
	return teams, nil
}
