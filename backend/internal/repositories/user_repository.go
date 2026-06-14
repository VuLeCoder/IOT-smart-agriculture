package repositories

import (
	"IOT-Smart-Agriculture/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

type userRepo struct {
	db *pgxpool.Pool
}

func CreateNewUserRepo(db *pgxpool.Pool) IUserRepository {
	return &userRepo{
		db: db,
	}
}

const (
	CREATE_USER_QUERY = `
		insert into users (id, email, password_hash, created_at)
		values ($1, $2, $3, $4)
	`

	GET_USER_BY_EMAIL_QUERY = `
		select
			id,
			email,
			password_hash 
		from users 
		where email = $1
	`
)

func (r *userRepo) CreateUser(ctx context.Context, user models.User) error {

	_, err := r.db.Exec(
		ctx, CREATE_USER_QUERY,
		user.ID, user.Email, user.PasswordHash, user.CreatedAt,
	)
	return err
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(ctx, GET_USER_BY_EMAIL_QUERY, email).Scan(&user.ID, &user.Email, &user.PasswordHash)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
