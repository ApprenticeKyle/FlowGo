package dto

import (
	"FLOWGO/pkg/utils"
)

// CreateProjectRequest 创建项目请求
type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	OwnerId     uint64 `json:"owner_id" binding:"required"`
}

// CreateProjectResponse 创建项目响应
type CreateProjectResponse struct {
	ID uint64 `json:"id"`
}

// UpdateProjectRequest 更新项目请求
type UpdateProjectRequest struct {
	ID          uint64     `json:"id" binding:"required"`
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description" binding:"required"`
	OwnerId     uint64     `json:"ownerId" binding:"required"`
	Status      string     `json:"status" binding:"required"`
	Deadline    utils.Time `json:"deadline" binding:"required"`
	StartDate   utils.Time `json:"startDate" binding:"required"`
	Members     int        `json:"members" binding:"required"`
	Progress    int        `json:"progress" binding:"required"`
	Priority    string     `json:"priority" binding:"required"`
	Starred     bool       `json:"starred" binding:"required"`
	Archived    bool       `json:"archived" binding:"required"`
	CoverImage  string     `json:"coverImage" binding:"omitempty"`
	TeamIds     []uint64   `json:"teamIds" binding:"omitempty"`
	Tags        []string   `json:"tags" binding:"omitempty"`
}

// UpdateProjectResponse 更新项目响应
type UpdateProjectResponse struct {
	ID          uint64     `json:"id"`
	TeamIds     []uint64   `json:"teamIds"`
	Tags        []string   `json:"tags"`
	Members     int        `json:"members"`
	Progress    int        `json:"progress"`
	Priority    string     `json:"priority"`
	Starred     bool       `json:"starred"`
	Archived    bool       `json:"archived"`
	CoverImage  string     `json:"coverImage"`
	Deadline    utils.Time `json:"deadline"`
	StartDate   utils.Time `json:"startDate"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	OwnerId     uint64     `json:"ownerId"`
	Status      string     `json:"status"`
}

// DeleteProjectRequest 删除项目请求
type DeleteProjectRequest struct {
	ID uint64 `json:"id" binding:"required"`
}

// DeleteProjectResponse 删除项目响应
type DeleteProjectResponse struct {
	ID uint64 `json:"id"`
}

// GetProjectRequest 获取项目请求
type GetProjectRequest struct {
	ID uint64 `uri:"id" binding:"required"`
}

// GetProjectResponse 获取项目响应
type GetProjectResponse struct {
	ID          uint64     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	OwnerId     uint64     `json:"owner_id"`
	Status      string     `json:"status"`
	Deadline    utils.Time `json:"deadline"`
	StartDate   utils.Time `json:"start_date"`
	Members     int        `json:"members"`
	Progress    int        `json:"progress"`
	Priority    string     `json:"priority"`
	Starred     bool       `json:"starred"`
	Archived    bool       `json:"archived"`
	CoverImage  string     `json:"cover_image"`
}

// ProjectListResponse 项目列表响应
type ProjectListResponse struct {
	List []*ProjectResponse `json:"list"`
	Page PageResponse       `json:"page"`
}

// ProjectResponse 项目响应
type ProjectResponse struct {
	ID          uint64     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	OwnerId     uint64     `json:"owner_id"`
	Status      string     `json:"status"`
	Deadline    utils.Time `json:"deadline"`
	StartDate   utils.Time `json:"start_date"`
	Members     int        `json:"members"`
	Progress    int        `json:"progress"`
	Priority    string     `json:"priority"`
}

// ProjectTeamsResponse 项目团队响应
type ProjectTeamsResponse struct {
	Teams []*TeamResponse `json:"teams"`
}

// TeamResponse 团队响应
type TeamResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
