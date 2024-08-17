package handlers

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"server/database/delete"
	"server/database/get"
	"server/database/insert"
	"server/models"
)

func GetPersonalProjects(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)

	result, err := get.ProjectsQuery(database, "personal_projects")
	if err != nil {
		log.Error().Err(err).Msg("[Database] Error getting personal projects.")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, result)
}
func GetTeamProjects(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)

	result, err := get.ProjectsQuery(database, "team_projects")
	if err != nil {
		log.Error().Err(err).Msg("[Database] Error getting team projects.")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusOK, result)
}

func InsertPersonalProjects(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	project := new(models.Project)

	err := c.Bind(project)
	if err != nil {
		log.Error().Err(err).Msg("[Database] Error parsing request body.")
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	err = insert.ProjectsInsert(database, *project, "personal_projects")
	if err != nil {
		log.Error().Err(err).Msg("[Database] Error parsing request body.")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(http.StatusCreated, project)
}

func DeletePersonalProjects(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	deleteValue := c.Param("id")

	err := delete.ProjectsDelete(database, deleteValue, "personal_projects")
	if err != nil {
		log.Error().Err(err).Msg("[Database] Error deleting personal projects.")
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.String(http.StatusAccepted, "Deleted")
}
