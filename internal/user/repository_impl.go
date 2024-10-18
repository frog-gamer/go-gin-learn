// internal/user/repository_impl.go
package user

import (
	"database/sql"
	"log"
	"myapp/pkg/pagination"
)

type pgRepo struct {
	db *sql.DB
}

// NewPGRepository creates a new instance of pgRepo which implements the Repository interface
func NewPGRepository(db *sql.DB) Repository {
	return &pgRepo{db: db}
}

// Create inserts a new user into the database
func (r *pgRepo) Create(user *User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2)`
	_, err := r.db.Exec(query, user.Email, user.Password)
	if err != nil {
		log.Println("Error executing SQL query:", err) // Log the SQL error
		return err
	}

	return nil
}

// GetByEmail retrieves a user by email from the database
func (r *pgRepo) GetByEmail(email string) (*User, error) {
	user := &User{}
	query := `SELECT id, email, password FROM users WHERE email=$1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}

// GetPaginated retrieves a paginated list of users from the database
func (r *pgRepo) GetPaginated(page, limit int) ([]User, int, error) {
	offset := pagination.GetOffset(page, limit) // Use the global pagination function
	query := `SELECT id, email FROM users LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email); err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	// Count the total number of rows
	var total int
	countQuery := `SELECT COUNT(*) FROM users`
	err = r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
