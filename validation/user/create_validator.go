package user

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"time"
)

func CreateUserValidator(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string, ) bool {

	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Before(date) {
			return false
		}
	}
	return true
}
