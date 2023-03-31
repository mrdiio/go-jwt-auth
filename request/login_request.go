package request

type LoginRequest struct {
	Email    string `binding:"required,email" form:"email"`
	Password string `binding:"required" form:"password"`
}
