package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"name" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}
