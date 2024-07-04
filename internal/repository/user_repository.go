package repository

import (
	"github.com/sirupsen/logrus"
	"go-clean-arch/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) FindByEmail(db *gorm.DB, user *entity.User, email string) error {
	return db.Where("email = ?", email).First(user).Error
}

func (r *UserRepository) FindByCredential(db *gorm.DB, user *entity.User, email, password string) error {
	return db.Where("email = ? AND password = ?", email, password).First(user).Error
}
