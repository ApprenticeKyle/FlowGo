package entity

import "time"

type Project struct {
	BaseEntity
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	OwnerId     uint64    `json:"owner_id" gorm:"not null"`
	Status      string    `json:"status"`
	Deadline    time.Time `json:"deadline"`
	StartDate   time.Time `json:"start_date"`
	Members     int       `json:"members"`
	Progress    int       `json:"progress"`
	Priority    string    `json:"priority"`
	Starred     bool      `json:"starred"`
	Archived    bool      `json:"archived"`
	CoverImage  string    `json:"cover_image"`
}

func (Project) TableName() string {
	return "projects"
}

func (p *Project) IsActive() bool {
	return p.Status == "active" && !p.IsDeleted()
}
