package entity

type Team struct {
	BaseEntity
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	OwnerId     uint64 `json:"owner_id"`
}

func (Team) TableName() string {
	return "teams"
}
