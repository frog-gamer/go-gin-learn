package user

import (
	"errors"
	"log"
	"myapp/config"
	"myapp/pkg/crypto"
	"myapp/pkg/jwt"
	"myapp/pkg/pagination"
)

type useCase struct {
	repo Repository
}

// NewUseCase creates a new use case instance
func NewUseCase(repo Repository) UseCase {
	return &useCase{repo: repo}
}

// Register handles user registration by encrypting the password and saving the user.
func (u *useCase) Register(user *User) error {
	// Print the key length for debugging
	log.Println("JWT_SECRET length:", len(config.GetJWTSecretKey()))

	// Encrypt the password using AES
	encryptedPassword, err := crypto.EncryptAES(user.Password, config.GetJWTSecretKey())
	if err != nil {
		log.Println("Error encrypting password:", err) // Log error for debugging
		return err
	}

	user.Password = encryptedPassword

	// Save the user in the database
	err = u.repo.Create(user)
	if err != nil {
		log.Println("Error saving user in repository:", err) // Log error for debugging
		return err
	}

	return nil
}

// Login handles user login by decrypting the stored password, comparing it, and generating a JWT token.
func (u *useCase) Login(email, password string) (string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Decrypt the stored password
	decryptedPassword, err := crypto.DecryptAES(user.Password, config.GetJWTSecretKey())
	if err != nil {
		return "", errors.New("failed to decrypt password")
	}

	// Compare the provided password with the decrypted password
	if password != decryptedPassword {
		return "", errors.New("incorrect password")
	}

	// Generate JWT token if authentication is successful
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUsers retrieves a paginated list of users.
func (u *useCase) GetUsers(page, limit int) (pagination.PaginationData, error) {
	users, total, err := u.repo.GetPaginated(page, limit)
	if err != nil {
		return pagination.PaginationData{}, err
	}

	// Use the global pagination function to structure the paginated data
	paginatedData := pagination.Paginate(page, limit, total, users)
	return paginatedData, nil
}
