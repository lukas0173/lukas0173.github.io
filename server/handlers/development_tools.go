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

func GetDevelopmentTools(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	queryResult, err := get.DevelopmentToolsGet(database)
	if err != nil {
		log.Error().Err(err).Msg("[Database] Error getting development tools.")
		return c.String(http.StatusInternalServerError, "Iternal server error")
	}
	return c.JSON(http.StatusOK, queryResult)
}

func InsertDevelopmentTools(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	developmentTools := new(models.DevelopmentTools)

	err := c.Bind(developmentTools)
	if err != nil {
		log.Error().Err(err).Msg("[Database] Error parsing development tools.")
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	err = insert.DevelopmentToolsInsert(database, *developmentTools)
	if err != nil {
		log.Error().Err(err).Msg("[Database] Error inserting development tools.")
		return c.String(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusCreated, developmentTools)
}

func DeleteDevelopmentTools(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	deleteValue := c.Param("id")

	err := delete.DevelopmentToolsDelete(database, deleteValue)
	if err != nil {
		log.Error().Err(err).Msg("[Database] Error deleting development tools.")
		return c.String(http.StatusInternalServerError, "Internal server error")
	}
	return c.String(http.StatusOK, "Successfully deleted development tools.")
}
