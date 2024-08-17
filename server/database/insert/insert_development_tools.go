package insert

import (
	"context"
	"github.com/jackc/pgx/v5"
	"server/models"
	"strconv"
	"strings"
)

func DevelopmentToolsInsert(database *pgx.Conn, developmentTools models.DevelopmentTools) error {

	// This is the slice of the contents inside the VALUES() command, will be separated by the comma
	inputDevelopmentToolsDataSlice := []string{"'" + developmentTools.Field + "'", "'" + developmentTools.Description + "'", "'" + developmentTools.Style.TextColor + "'", strconv.Itoa(developmentTools.Style.Span), strconv.Itoa(developmentTools.Order)}

	// Inserts into two tables using WITH
	queryCommand := "WITH favourite_tools_insertion " +
		"AS (INSERT INTO favourite_tools (field, descriptions, text_color, span, \"order\") VALUES (" + strings.Join(inputDevelopmentToolsDataSlice, ",") + ") RETURNING field) " + "INSERT INTO favourite_tools_icons (field, path, icon_names)" + "VALUES ((SELECT field FROM favourite_tools_insertion), '" + developmentTools.Icons.Path + "', ARRAY['" + strings.Join(developmentTools.Icons.Name, "','") + "'])"

	_, err := database.Exec(context.Background(), queryCommand)
	if err != nil {
		return err
	}

	return nil
}
