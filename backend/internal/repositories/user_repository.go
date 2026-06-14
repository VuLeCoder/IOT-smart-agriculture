package repositories

import (
	"IOT-Smart-Agriculture/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) error
}

type userRepo struct {
	db *pgxpool.Pool
}

func CreateNewUserRepo(db *pgxpool.Pool) IUserRepository {
	return &userRepo{
		db: db,
	}
}

const ()

func (r *userRepo) CreateUser(ctx context.Context, user models.User) error {
	return nil
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) error {
	return nil
}
