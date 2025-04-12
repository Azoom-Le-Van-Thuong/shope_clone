package valueobject

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	SaltLength = 16
	BcryptCost = bcrypt.DefaultCost // Bạn có thể tăng lên nếu muốn
)

type Password struct {
	salt string
	hash string
}

// NewPasswordFromPlainText tạo password mới từ plaintext
func NewPasswordFromPlainText(ctx context.Context, plain string) (*Password, error) {
	salt, err := generateSalt(SaltLength)
	if err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(salt+plain), BcryptCost)
	if err != nil {
		return nil, err
	}

	return &Password{
		salt: salt,
		hash: string(hashed),
	}, nil
}

// NewPasswordFromHash tái tạo lại object từ DB
func NewPasswordFromHash(salt, hash string) *Password {
	return &Password{
		salt: salt,
		hash: hash,
	}
}

// Compare kiểm tra password người dùng nhập có đúng không
func (p *Password) Compare(ctx context.Context, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(p.salt+input))
}

// Hash trả về hash để lưu DB
func (p *Password) Hash() string {
	return p.hash
}

// Salt trả về salt để lưu DB
func (p *Password) Salt() string {
	return p.salt
}

// Tạo salt ngẫu nhiên
func generateSalt(length int) (string, error) {
	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}
	return base64.RawStdEncoding.EncodeToString(buf), nil
}
