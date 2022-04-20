package user

type User struct {
	Id       int    `json:"-"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
}
