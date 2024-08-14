package delete

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func ProjectsDelete(database *pgx.Conn, value, projectType string) error {
	_, err := database.Exec(context.Background(), "DELETE FROM "+projectType+" WHERE id = "+value)
	if err != nil {
		return err
	}

	return nil
}
