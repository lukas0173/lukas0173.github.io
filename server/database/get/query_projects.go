package get

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"server/models"
)

func ProjectsQuery(database *pgx.Conn, projectType string) []models.Project {
	// Query command with dynamic table
	query, err := database.Query(context.Background(), "SELECT id, project_name, project_description, link, technologies, image_name FROM "+projectType+" ORDER BY id")
	if err != nil {
		log.Err(err).Msg("[Database] Error querying personal projects: %v\n")
	}

	// Iterate through the get results and assign it into returned value
	var returnedProjects []models.Project
	for query.Next() {
		// Holding get values for each field
		var (
			id                 int
			projectName        string
			projectDescription string
			link               string
			technologies       []string
			imageName          string
		)
		err := query.Scan(&id, &projectName, &projectDescription, &link, &technologies, &imageName)
		if err != nil {
			log.Err(err).Msg("[Database] Error scanning personal projects: %v\n")
			return nil
		}

		returnedProjects = append(returnedProjects, models.Project{
			Id: id, ProjectName: projectName, ProjectDescription: projectDescription, Link: link, Technologies: technologies, ImageName: imageName,
		})
	}
	// Close get result on function return
	defer query.Close()

	return returnedProjects
}
