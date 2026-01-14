package entity

type ProjectsTeams struct {
	BaseEntity
	ProjectId uint64 `json:"project_id" gorm:"not null"`
	TeamId    uint64 `json:"team_id" gorm:"not null"`
}

func (ProjectsTeams) TableName() string {
	return "projects_teams"
}
