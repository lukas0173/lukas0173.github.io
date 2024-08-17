package delete

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func DevelopmentToolsDelete(database *pgx.Conn, value string) error {
	_, err := database.Exec(context.Background(), "DELETE FROM favourite_tools WHERE id = $1", value)
	if err != nil {
		return err
	}
	return nil
}
