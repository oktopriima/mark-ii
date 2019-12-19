package users

import (
	"time"
)

type CrateRequest struct {
	Name     string    `json:"name" form:"name" binding:"required"`
	Email    string    `json:"email" form:"email" binding:"email"`
	Birthday time.Time `json:"birthday" form:"birthday" binding:"required"`
	Password string    `json:"password" binding:"required" form:"password"`
}
