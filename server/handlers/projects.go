package handlers

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/database/query"
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
