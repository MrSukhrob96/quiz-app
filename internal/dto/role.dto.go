package dto

type CreateRoleDTO struct {
	Name string `json:"name"`
}

type UpdateRoleDTO struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
