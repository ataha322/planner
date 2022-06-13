package models

type User struct {
	ID        uint   `json: "user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm: "unqiue"`
	Password  []byte `json:"-"`
}
