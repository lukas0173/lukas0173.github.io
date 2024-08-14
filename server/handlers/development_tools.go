package handlers

import (
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"server/database/get"
)

func GetDevelopmentTools(c echo.Context) error {
	database := c.Get("database").(*pgx.Conn)
	queryResult := get.DevelopmentToolsQuery(database)
	return c.JSON(http.StatusOK, queryResult)
}
