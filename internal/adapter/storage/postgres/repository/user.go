package repository

import (
	"database/sql"
	"log/slog"

	"github.com/AmirHosein-Gharaati/user-management/internal/adapter/storage/postgres"
	"github.com/AmirHosein-Gharaati/user-management/internal/core/domain"
)

type UserRepositoryImpl struct {
	db *postgres.DB
}

func NewUserRepository(db *postgres.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) CreateUser(user *domain.User) (*domain.User, error) {
	row := r.db.DB.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *", user.Name, user.Email, user.Password)

	var userDB domain.User
	err := row.Scan(
		&userDB.ID,
		&userDB.Name,
		&userDB.Email,
		&userDB.Password,
		&userDB.CreatedAt,
	)
	if err != nil {
		slog.Error("error while scanning the user from db", "error", err)
		return nil, err
	}

	return &userDB, nil
}

func (r *UserRepositoryImpl) ExistsUserByEmail(email string) bool {
	var exists int
	err := r.db.DB.QueryRow("SELECT 1 FROM users WHERE email=$1 LIMIT 1", email).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		slog.Error("error querying for user existence by email", "error", err)
		return false
	}
	return true
}
