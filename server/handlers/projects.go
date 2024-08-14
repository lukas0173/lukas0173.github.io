package handlers

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/database/insert"
	"server/database/query"
	"server/models"
)

func GetPersonalProjects(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	result := query.ProjectsQuery(database, "personal_projects")
	return c.JSON(http.StatusOK, result)
}
func GetTeamProjects(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	result := query.ProjectsQuery(database, "team_projects")
	return c.JSON(http.StatusOK, result)
}

func InsertPersonalProjects(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	project := new(models.Project)

	err := c.Bind(project)
	if err != nil {
		return err
	}

	err = insert.ProjectsInsert(database, *project, "personal_projects")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, project)
}
