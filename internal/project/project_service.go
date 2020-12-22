package project

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
)

func getProjects(db *gorm.DB) (GetProjectsResponseDTO, int, error) {
	var projects []Project

	result := db.Find(&projects)

	if result.Error != nil {
		return GetProjectsResponseDTO{}, http.StatusNotFound, errors.New("Project not found")
	}

	return createProjectsDTO(&projects), http.StatusOK, nil
}

func createProjectsDTO(projects *[]Project) GetProjectsResponseDTO {
	var result GetProjectsResponseDTO

	var projectsDTO []GetProjectResponseDTO
	for _, project := range *projects {
		projectsDTO = append(projectsDTO, createProjectDTO(&project))
	}

	result.Projects = projectsDTO
	return result
}

func createProjectDTO(project *Project) GetProjectResponseDTO {
	return GetProjectResponseDTO{
		UUID:        project.UUID,
		Name:        project.Name,
		Description: project.Description,
	}
}

// Update project
// Post project
// Delete project
