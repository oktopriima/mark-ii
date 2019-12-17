package services

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-ii/model"
)

type UserServiceContract interface {
	Create(user *model.User, db *gorm.DB) error
}

type userServiceContractService struct {
	db *gorm.DB
}

func NewUserServiceContract(db *gorm.DB) UserServiceContract {
	return &userServiceContractService{db}
}

func (srv *userServiceContractService) Create(user *model.User, tx *gorm.DB) error {
	var err error
	err = tx.Create(&user).Error
	return err
}
