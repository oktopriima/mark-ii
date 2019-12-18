package services

import (
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/mark-ii/model"
	"time"
)

type UserServiceContract interface {
	Create(user *model.User, tx *gorm.DB) error
	Update(user *model.User, tx *gorm.DB) error
	Find(ID int) (*model.User, error)
	FindBy(criteria map[string]interface{}) ([]*model.User, error)
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
	err = tx.Save(&user).Error

	return err
}

func (srv *userContractService) Find(ID int) (*model.User, error) {
	user := new(model.User)
	var err error

	err = srv.db.Where("id=?", ID).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (srv *userContractService) FindBy(criteria map[string]interface{}) ([]*model.User, error) {
	user := []*model.User{}
	var err error

	err = srv.db.Where(criteria).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
