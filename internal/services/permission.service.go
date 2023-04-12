package services

type PermissionService interface {
}

type permissionService struct {
}

func NewPermissionService() PermissionService {
	return &permissionService{}
}
