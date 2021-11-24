package types

type User struct {
	Id       string `json:"-" db:"id"`
	Name     string `json:"name"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
