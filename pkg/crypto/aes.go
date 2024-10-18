package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// PKCS7Padding pads the input to be a multiple of the block size
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS7UnPadding removes the padding after decryption
func PKCS7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("decryption error: input too short")
	}
	padding := int(data[length-1])
	if padding > length {
		return nil, errors.New("decryption error: invalid padding")
	}
	return data[:length-padding], nil
}

// EncryptAES encrypts the given data using AES-256-CBC encryption with PKCS7 padding
func EncryptAES(data, key string) (string, error) {
	// Ensure the key is 16, 24, or 32 bytes long
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", aes.KeySizeError(len(key)) // Return an error if the key size is invalid
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	plaintext := PKCS7Padding([]byte(data), block.BlockSize()) // Apply PKCS7 padding

	// Create initialization vector (IV)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// Return base64 encoded ciphertext
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAES decrypts the given base64 encoded AES-256-CBC encrypted data with PKCS7 unpadding
func DecryptAES(data, key string) (string, error) {
	// Ensure the key is 16, 24, or 32 bytes long
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", aes.KeySizeError(len(key)) // Return an error if the key size is invalid
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Remove padding after decryption
	plaintext, err := PKCS7UnPadding(ciphertext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
