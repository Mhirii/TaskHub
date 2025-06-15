package main

import (
	"flag"

	"github.com/Mhirii/TaskHub/backend/handlers"
	"github.com/Mhirii/TaskHub/backend/pkg/config"
	"github.com/Mhirii/TaskHub/backend/pkg/infra"
	"github.com/Mhirii/TaskHub/backend/repo"
	"github.com/Mhirii/TaskHub/backend/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load Config
	configPath := flag.String("config", "etc/config.yaml", "path to config file")
	flag.Parse()
	cfg, err := config.ParseConfig(*configPath)
	if err != nil {
		panic(err)
	}

	// DB
	db, err := infra.InitDB(*cfg)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	authRepo := repo.NewUsersRepo(db)
	authSvc := services.NewAuthService(authRepo)

	tasksRepo := repo.NewTasksRepo(db)
	tasksSvc := services.NewTasksService(tasksRepo)

	projectsRepo := repo.NewProjectsRepo(db)
	projectsSvc := services.NewProjectsService(projectsRepo)

	usersRepo := repo.NewUsersRepo(db)
	usersSvc := services.NewUsersService(usersRepo)

	handlers.NewAuthHandlers(authSvc).WriteGroup(e.Group("/auth"))
	handlers.NewBoardHandlers().WriteGroup(e.Group("/boards"))
	handlers.NewProjectHandlers(projectsSvc).WriteGroup(e.Group("/projects"))
	handlers.NewTaskHandlers(tasksSvc).WriteGroup(e.Group("/tasks"))
	handlers.NewUsersHandlers(usersSvc).WriteGroup(e.Group("/users"))

	e.Logger.Fatal(e.Start(":" + cfg.Server.Port))
}
