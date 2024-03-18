package models

type User struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password"`
	Role     string `json:"role" binding:"required" db:"role"`
	RoleId   int    `json:"role_id" binding:"required" db:"role_id"`
}

type Role struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
