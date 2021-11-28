package dto

type LoginDTO struct{
	Email string `binding:"required" form:"email" json:"email" validate:"email"`
	Password string `binding:"required" form:"password" json:"password" validate:"min:8"`
}