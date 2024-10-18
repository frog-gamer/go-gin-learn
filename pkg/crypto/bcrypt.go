
package crypto

import "golang.org/x/crypto/bcrypt"

// EncryptPassword hashes the password using bcrypt
func EncryptPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPasswordHash checks password hash using bcrypt
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
    