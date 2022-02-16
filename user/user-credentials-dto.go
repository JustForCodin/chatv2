package user

type UserCredentialsDTO struct {
	Email    string `gorm:"unique;not_null"`
	Password string // # bcrypt
}
