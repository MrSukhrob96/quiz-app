package repositories

import (
	"quiz-app/internal/dto"
	"quiz-app/internal/entities"

	uuid "github.com/satori/go.uuid"
)

type RoleRepository interface {
	GetAll() (entities.Roles, error)
	GetById(id *uuid.UUID) (entities.Role, error)
	Create(userDTO dto.CreateRoleDTO) (entities.Role, error)
	Update(id *uuid.UUID, userDTO dto.UpdateRoleDTO) (entities.Role, error)
	Delete(id *uuid.UUID) error
}

type roleRepository struct {
}

func NewRoleRepository() RoleRepository {
	return &roleRepository{}
}

func (repository *roleRepository) GetAll() (entities.Roles, error) {
	return nil, nil
}

func (repository *roleRepository) GetById(id *uuid.UUID) (entities.Role, error) {
	return entities.Role{}, nil
}

func (repository *roleRepository) Create(userDTO dto.CreateRoleDTO) (entities.Role, error) {
	return entities.Role{}, nil
}

func (repository *roleRepository) Update(id *uuid.UUID, userDTO dto.UpdateRoleDTO) (entities.Role, error) {
	return entities.Role{}, nil
}

func (repository *roleRepository) Delete(id *uuid.UUID) error {
	return nil
}
