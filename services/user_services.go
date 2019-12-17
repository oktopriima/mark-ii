package services

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-ii/model"
	"time"
)

type UserServiceContract interface {
	Create(user *model.User, tx *gorm.DB) error
	Update(user *model.User, tx *gorm.DB) error
}

type userContractService struct {
	db *gorm.DB
}

func NewUserServiceContract(db *gorm.DB) UserServiceContract {
	return &userContractService{db}
}

func (srv *userContractService) Create(user *model.User, tx *gorm.DB) error {
	var err error
	// err = tx.Create(&user).Error
	query := "INSERT INTO pengguna (name, email, password, created_at, updated_at) VALUES (?,?,?,?,?)"
	err = tx.Exec(query, user.Name, user.Email, "XXX", time.Now(), time.Now()).Error
	return err
}

func (srv *userContractService) Update(user *model.User, tx *gorm.DB) error {
	var err error

	// TO DO update query
	err = tx.Update(&user).Error

	return err
}
