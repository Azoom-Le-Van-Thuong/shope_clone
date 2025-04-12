package valueobject

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	salt   string
	hashed string
}

func generateSalt(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func NewPassword(raw string, saltRounds int) (Password, error) {
	if len(raw) < 6 {
		return Password{}, errors.New("password too short")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, err
	}
	salt, err := generateSalt(saltRounds)
	if err != nil {
		return Password{}, err
	}
	return Password{hashed: string(hash), salt: salt}, nil
}

func FromHashed(hash string) Password {
	return Password{hashed: hash}
}

func (p Password) Verify(raw string, salt string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.hashed), []byte(raw+salt)) == nil
}

func (p Password) Hashed() string {
	return p.hashed
}
func (p Password) Salt() string {
	return p.salt
}
