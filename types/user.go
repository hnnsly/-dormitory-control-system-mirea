package types

type User struct {
	Id       uint   `json:"id"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}
