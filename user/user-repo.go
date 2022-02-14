package user

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserFilter struct {
	ChatID   *int64
	UserName *string
}

type UserRepository interface {
	GetUser(userID, initiatorID int64) (*User, error)
	GetUsers(userID int64, filter *UserFilter) ([]User, error)
	CreateUser(userID int64, user User) (*User, error)
	UpdateUser(userID int64, user User) (*User, error)
	DeleteUser(userID, initiatorID int64) (*User, error)
}

type UserRepositoryImpl struct {
	db  *gorm.DB
	err error
}

func (r *UserRepositoryImpl) GetUser(userID, initiatorID int64) (*User, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}

	var user User
	r.db.Where("id=?", userID).Find(&user)
	return &user, r.err
}

func (r *UserRepositoryImpl) GetUsers(userID int64, filter *UserFilter) ([]User, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}

	var users []User
	r.db.Where("id=?, UserName=?", filter.ChatID, filter.UserName).Find(&users)
	return users, r.err
}

func (r *UserRepositoryImpl) CreateUser(userID int64, user User) (*User, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}
	r.db.Create(user)
	return &user, r.err
}

func (r *UserRepositoryImpl) UpdateUser(userID int64, user User) (*User, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}

	var userUpdate User
	r.db.Where("text = ?", userUpdate.Name).Find(&userUpdate)
	user.Name = userUpdate.Name
	r.db.Save(&userUpdate)
	return &userUpdate, r.err
}

func (r *UserRepositoryImpl) DeleteUser(userID, initiatorID int64) (*User, error) {
	r.db, r.err = gorm.Open(mysql.Open("text"), &gorm.Config{})
	if r.err != nil {
		log.Fatal(r.err)
	}

	var user User
	r.db.Where("name = ?", user.Name).Find(&user)
	r.db.Delete(&user)
	return &user, r.err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}
