package token

type Token interface {
	GenerateToken(userID uint, email string) (string, error)
}
