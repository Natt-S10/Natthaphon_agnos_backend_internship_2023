package repository

import (
	"context"

	"github.com/Natt-S10/Natthaphon_agnos_backend_internship_2023/srcs/domain/models"
)

// Repository represent the todo repository contranct
type Repository interface {
	Create(ctx context.Context, data *models.ToDo) (int, error)
	Read(ctx context.Context, id int) (*models.ToDo, error)
}

type systemLog struct {
}
