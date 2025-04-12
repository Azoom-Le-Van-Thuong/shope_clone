package valueobject

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	regex := `^[a-zA-Z0-9._%%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	if !regexp.MustCompile(regex).MatchString(value) {
		return Email{}, errors.New("invalid email format")
	}
	return Email{value: value}, nil
}

func (e Email) Value() string {
	return e.value
}
