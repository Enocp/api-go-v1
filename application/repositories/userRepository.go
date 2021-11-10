package repositories

import "github.com/Enocp/desafios/domain"

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}