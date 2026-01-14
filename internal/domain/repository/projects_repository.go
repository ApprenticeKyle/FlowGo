package repository

import (
	"FLOWGO/internal/domain/entity"
)

type ProjectsRepository interface {
	BaseRepository[entity.Project]
}
