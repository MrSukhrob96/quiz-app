package dto

import uuid "github.com/google/uuid"

type CreateUserDTO struct {
	Name        string     `json:"name"`
	Branch      *uuid.UUID `json:"branch"`
	Office      *uuid.UUID `json:"office"`
	PhoneNumber string     `json:"phoneNumber"`
}

type UpdateUserDTO struct {
	Name        string     `json:"name"`
	Branch      *uuid.UUID `json:"branch"`
	Office      *uuid.UUID `json:"office"`
	PhoneNumber string     `json:"phoneNumber"`
	Status      string     `json:"status"`
}
