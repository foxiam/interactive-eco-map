package repository

import (
	"context"

	"user-microservice/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	const sql = "SELECT * FROM public.user WHERE id = $1"

	var user model.User
	err := r.db.QueryRow(ctx, sql, id).Scan(&user.ID, &user.Email, &user.Password)
	return &user, err
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	const sql = "SELECT * FROM public.user"

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	const sql = "SELECT * FROM public.user WHERE email = $1"

	var user model.User
	err := r.db.QueryRow(ctx, sql, email).Scan(&user.ID, &user.Email, &user.Password)
	return &user, err
}

func (r *UserRepository) AddUser(user *model.User) error {
	const sql = "INSERT INTO public.user(email, password) VALUES ($1, $2)"
	_, err := r.db.Exec(context.Background(), sql, user.Email, user.Password)
	return err
}

func (r *UserRepository) DeleteUser(id string) error {
	const sql = "DELETE FROM public.user WHERE id = $1"
	_, err := r.db.Exec(context.Background(), sql, id)
	return err
}