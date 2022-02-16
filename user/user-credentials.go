package user

type UserCredentials struct {
	Email    string `gorm:"unique;not_null"`
	Password string
}
