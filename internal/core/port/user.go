package port

import "github.com/AmirHosein-Gharaati/user-management/internal/core/domain"

// UserRepository is an interface for interacting with user-related DATA
type UserRepository interface {
	CreateUser(user *domain.User) (*domain.User, error)
	ExistsUserByEmail(email string) bool
}

// UserService is an interface for interacting with user-related BUSINESS LOGIC
type UserService interface {
	Register(user *domain.User) (*domain.User, error)
}
