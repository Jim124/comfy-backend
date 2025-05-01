package repository

import (
	"context"

	"github.com/jim124/comfy-backend/models"
)

type UserRepository interface {
	Close()
	Register(ctx context.Context, identifier string, password string) (*models.User, error)
	Login(ctx context.Context, identifier string, password string) (string, error)
}
