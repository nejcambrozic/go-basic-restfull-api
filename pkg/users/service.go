package users

import (
	"github.com/jinzhu/gorm"
	"github.com/nejcambrozic/go-basic-restfull-api/pkg/models"
)

type Service interface {
	GetUsers() []models.User
	CreateUser(user models.User)
	GetUserById(id int64) models.User
	UpdateUserById(id int64, user models.User)
	DeleteUserById(id int64) models.User
}

type service struct {
	db gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: *db}
}

func (s *service) CreateUser(user models.User) {
	// TODO handle duplicated email, username, etc
	s.db.Create(&user)
	return
}

func (s *service) GetUsers() []models.User {
	// TODO pagination
	var users []models.User
	s.db.Find(&users)
	return users
}

func (s *service) GetUserById(id int64) models.User {
	// TODO: handle user not found
	var user models.User
	s.db.First(&user, id)
	return user
}

func (s *service) UpdateUserById(id int64, user models.User){
	// TODO: handle user not found, validation
	var dbUser models.User
	s.db.First(&dbUser, id)
	// TODO: there must be a better way to do this
	s.db.Model(&dbUser).Update("email", user.Email)
	s.db.Model(&dbUser).Update("username", user.Username)
}


func (s *service) DeleteUserById(id int64) models.User {
	var user models.User
	s.db.First(&user, id)
	s.db.Delete(&user)
	return user
}