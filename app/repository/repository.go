package repository

import (
	"golang.org/x/net/context"
	"wiki/app/models"
)

type PostRepo interface {
	Fetch(ctx context.Context, num int64)([]*models.Post, error)
	GetByID(ctx context.Context, id int64)(*models.Post, error)
	Create(ctx context.Context, p *models.Post) (int64, error)
	Update(ctx context.Context, p *models.Post) (*models.Post, error)
	Delete(ctx context.Context, id int64) (bool, error)
}