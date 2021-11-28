package dto

//UserUpdateDTO used when PUT method is impelement
type UserUpdateDTO struct {
	ID       uint64 `binding:"required" form:"id" json:"id"`
	Name     string `binding:"required" form:"name" json:"name"`
	Email    string `binding:"required" form:"email" json:"email"`
	Password string `binding:"required" form:"password,omitempty" json:"password,omitempty" validate:"min:8"`
}


//UserCreateDTO used when we want to create new user
type UserCreateDTO struct {
	ID       uint64 `binding:"required" form:"id" json:"id"`
	Name     string `binding:"required" form:"name" json:"name"`
	Email    string `binding:"required" form:"email" json:"email" validate:"email"`
	Password string `binding:"required" form:"password,omitempty" json:"password,omitempty" validate:"min:8"`
}