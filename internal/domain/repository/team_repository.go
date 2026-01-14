package repository

import (
	"FLOWGO/internal/domain/entity"
	"context"
)

type TeamRepository interface {
	ListAvailableTeams(ctx context.Context) ([]*entity.Team, error)
}
