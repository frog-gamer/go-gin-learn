package user

type Repository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	GetPaginated(page, limit int) ([]User, int, error)
}
