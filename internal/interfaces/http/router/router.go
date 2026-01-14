package router

import (
	"FLOWGO/internal/application/usecase"
	"FLOWGO/internal/infrastructure/database"
	"FLOWGO/internal/infrastructure/repository"
	"FLOWGO/internal/interfaces/http/handler"
	"FLOWGO/internal/interfaces/http/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.New()

	// 中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 依赖注入
	// 用户相关
	userRepo := repository.NewUserRepository(database.DB)
	createUserUseCase := usecase.NewCreateUserUseCase(userRepo)
	getUserUseCase := usecase.NewGetUserUseCase(userRepo)
	listUsersUseCase := usecase.NewListUsersUseCase(userRepo)
	loginUseCase := usecase.NewLoginUseCase(userRepo)

	// 项目相关
	projectRepo := repository.NewProjectsRepository(database.DB)
	createProjectUseCase := usecase.NewCreateProjectsCase(projectRepo)
	updateProjectUseCase := usecase.NewUpdateProjectsCase(projectRepo)
	deleteProjectUseCase := usecase.NewDeleteProjectsCase(projectRepo)
	getProjectUseCase := usecase.NewGetProjectsCase(projectRepo)
	listProjectsUseCase := usecase.NewListProjectsCase(projectRepo)

	// 团队相关
	teamRepo := repository.NewTeamRepository(database.DB)
	projectsTeamUseCase := usecase.NewProjectsTeamUseCase(teamRepo)

	projectHandler := handler.NewProjectsHandler(createProjectUseCase, updateProjectUseCase, getProjectUseCase, listProjectsUseCase, deleteProjectUseCase, projectsTeamUseCase)
	userHandler := handler.NewUserHandler(createUserUseCase, getUserUseCase, listUsersUseCase)
	authHandler := handler.NewAuthHandler(loginUseCase)

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 认证相关路由（无需认证）
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}

		// 项目相关路由
		projects := v1.Group("/projects")
		projects.Use(middleware.Auth())
		{
			projects.GET("", projectHandler.ListProjects)
			projects.POST("", projectHandler.CreateProject)
			projects.GET("/:id", projectHandler.GetProject)
			projects.PUT("/:id", projectHandler.UpdateProject)
			projects.DELETE("/:id", projectHandler.DeleteProject)
			projects.GET("/teams/available", projectHandler.ProjectTeams)
		}

		// 用户相关路由
		users := v1.Group("/users")
		users.Use(middleware.Auth())
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.ListUsers)
			users.GET("/:id", userHandler.GetUser)
		}
	}

	return r
}
