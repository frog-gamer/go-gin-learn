package user

import "myapp/pkg/pagination"

type UseCase interface {
	Register(user *User) error
	Login(email, password string) (string, error)
	GetUsers(page, limit int) (pagination.PaginationData, error)
}
