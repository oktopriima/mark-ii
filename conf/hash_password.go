package conf

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	userPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(userPassword), err
}
