package requests

type CreateUserRequest struct {
	FirstName            string `form:"first_name" json:"first_name" binding:"required,min=3,max=100"`
	LastName             string `form:"last_name" json:"last_name" binding:"required,min=3,max=100"`
	Username             string `form:"username" json:"username" binding:"required,min=3,max=100"`
	Email                string `form:"email" json:"email" binding:"required,min=3,max=100"`
	Password             string `form:"password" json:"password" binding:"required,min=3,max=100"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required,min=3,max=100"`
}
