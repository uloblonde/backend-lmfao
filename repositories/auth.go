package repositories

import (
	"backendtask/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(email string) (models.User, error)
	CheckAuth(Id int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repo {
	return &repo{db}
}

func (r *repo) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
func (r *repo) Login(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}
func (r *repo) CheckAuth(Id int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, Id).Error
	return user, err
}
