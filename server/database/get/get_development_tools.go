package get

import (
	"context"
	"github.com/jackc/pgx/v5"
	"server/models"
)

func DevelopmentToolsGet(database *pgx.Conn) ([]models.DevelopmentTools, error) {
	// Query command
	query, err := database.Query(context.Background(), "SELECT tools.id, tools.field, tools.descriptions, tools.text_color, tools.span, tools.order, icons.path, icons.icon_names FROM favourite_tools as tools INNER JOIN favourite_tools_icons as icons on tools.id= icons.id ORDER BY tools.order")
	if err != nil {
		return nil, err
	}

	// Iterate through the get results and assign it into returned value
	var returnedLiteral []models.DevelopmentTools
	for query.Next() {
		// Holding get values for each field
		var field, description, textColor, path string
		var iconNames []string
		var id, span, order int

		err := query.Scan(&id, &field, &description, &textColor, &span, &order, &path, &iconNames)
		if err != nil {
			return nil, err
		}

		returnedLiteral = append(returnedLiteral, models.DevelopmentTools{
			Id:          id,
			Field:       field,
			Description: description,
			Style:       models.Style{Span: span, TextColor: textColor},
			Icons:       models.Icons{Path: path, Name: iconNames},
			Order:       order,
		})
	}

	// Close get result on function return
	defer query.Close()

	return returnedLiteral, nil
}
