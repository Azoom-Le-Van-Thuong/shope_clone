package domain

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name" gorm:"name"`
	Password string `json:"password" gorm:"password"`
	Salt     string `json:"salt" gorm:"salt"`
	Email    string `json:"email" gorm:"email uniqueIndex"`
}
