package entity

type User struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Id        int64  `json:"id"`
}
