package user

type UserCredentialsDTO struct {
	Email string // # unique
	Password string // # bcrypt
}