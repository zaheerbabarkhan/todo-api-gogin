package user

type CreateUserRequest struct {
	FirstName string `json:"firstName" binding:"required,min=1,max=20"`
	LastName  string `json:"lastName" binding:"required,min=1,max=20"`
	Email     string `json:"email" binding:"required,email,max=250"`
	Password  string `json:"password" binding:"required,min=8,max=15"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email,max=250"`
	Password string `json:"password" binding:"required,min=8,max=15"`
}
