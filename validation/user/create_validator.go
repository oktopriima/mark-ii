package user

import (
	"gopkg.in/go-playground/validator.v9"
	"time"
)

var CreateUserValidator validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.Before(date) {
			return false
		}
	}
	return true
}
