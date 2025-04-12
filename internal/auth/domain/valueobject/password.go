package valueobject

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hashed string
}

func NewPassword(raw string) (Password, error) {
	if len(raw) < 6 {
		return Password{}, errors.New("password too short")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, err
	}
	return Password{hashed: string(hash)}, nil
}

func FromHashed(hash string) Password {
	return Password{hashed: hash}
}

func (p Password) Verify(raw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p.hashed), []byte(raw)) == nil
}

func (p Password) Hashed() string {
	return p.hashed
}
