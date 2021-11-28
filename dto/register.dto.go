package dto

//RegisterDTO used when a post request to /register
type RegisterDTO struct {
	Name     string `binding:"required" form:"name" json:"name" validate:"min:1"`
	Email    string `binding:"required" form:"email" json:"email" validate:"email"`
	Password string `binding:"required" form:"password" json:"password" validate:"min:8"`
}
