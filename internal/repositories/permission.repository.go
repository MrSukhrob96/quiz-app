package repositories

import (
	"quiz-app/internal/dto"
	"quiz-app/internal/entities"

	uuid "github.com/google/uuid"
)

type PermissionRepository interface {
	GetAll() (entities.Permissions, error)
	GetById(id *uuid.UUID) (entities.Permission, error)
	Create(userDTO dto.CreatePermissionDTO) (entities.Permission, error)
	Update(id *uuid.UUID, userDTO dto.UpdatePermissionDTO) (entities.Permission, error)
	Delete(id *uuid.UUID) error
}

type permissionRepository struct {
}

func NewPermissionRepository() PermissionRepository {
	return &permissionRepository{}
}

func (repository *permissionRepository) GetAll() (entities.Permissions, error) {
	return nil, nil
}

func (repository *permissionRepository) GetById(id *uuid.UUID) (entities.Permission, error) {
	return entities.Permission{}, nil
}

func (repository *permissionRepository) Create(userDTO dto.CreatePermissionDTO) (entities.Permission, error) {
	return entities.Permission{}, nil
}

func (repository *permissionRepository) Update(id *uuid.UUID, userDTO dto.UpdatePermissionDTO) (entities.Permission, error) {
	return entities.Permission{}, nil
}

func (repository *permissionRepository) Delete(id *uuid.UUID) error {
	return nil
}
