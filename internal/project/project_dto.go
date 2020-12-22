package project

type GetProjectResponseDTO struct {
	UUID        string `json:"projectId"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type GetProjectsResponseDTO struct {
	Projects []GetProjectResponseDTO `json:"projects"`
}
