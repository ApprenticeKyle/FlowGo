package repository

import (
	"FLOWGO/internal/domain/entity"
	"context"
)

type ProjectsTeamsRepository interface {
	DeleteByProjectId(ctx context.Context, projectId uint64) error
	SaveBatch(ctx context.Context, entities []*entity.ProjectsTeams) error
}
