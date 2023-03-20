package request

type UserRequest struct {
	Name     string  `binding:"required,min=3,max=50" form:"name"`
	Username string  `binding:"required,min=3,max=50" form:"username"`
	Email    *string `binding:"omitempty,email" form:"email"`
	Password string  `binding:"required,min=6,max=50" form:"password"`
}
