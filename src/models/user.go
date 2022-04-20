package models

type User struct {
	Id        uint     `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"flast_name"`
	Email     string   `json:"email" gorm: "unqiue"`
	Password  []byte   `json:"-"`
	TaskList  taskList `json:"-"`
}
