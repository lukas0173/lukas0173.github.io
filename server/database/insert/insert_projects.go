package insert

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"server/models"
	"strings"
)

func ProjectsInsert(database *pgx.Conn, project models.Project, projectType string) error {
	inputDataSlice := []string{project.ProjectName, project.ProjectDescription, project.ImageName, project.Link}
	inputData := strings.Join(inputDataSlice, "','")

	queryCommand := "INSERT INTO " + projectType + " (project_name, project_description, image_name, link, technologies) VALUES ('" + inputData + "'," + "ARRAY['" + strings.Join(project.Technologies, "','") + "'])"

	_, err := database.Exec(context.Background(), queryCommand)
	if err != nil {
		log.Error().Err(err).Msg("[Database] Error inserting project")
		return err
	}

	return nil
}
