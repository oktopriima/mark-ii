package users

import "time"

type CrateRequest struct {
	Name     string    `json:"name" form:"name" binding:"required"`
	Email    string    `json:"email" form:"email"`
	Birthday time.Time `json:"birthday" form:"birthday" binding:"required,createuservalidator"`
}
