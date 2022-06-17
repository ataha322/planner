package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	Tasks     []Task `json:"tasks"` //user will contain his unique slice of tasks
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (user *User) AddTask(newTask Task) {
	user.Tasks = append(user.Tasks, newTask)
}

func (user *User) DeleteTask(deleteTask Task) {
	for index, task := range user.Tasks {
		if task.Id == deleteTask.Id {
			user.Tasks = append(user.Tasks[:index], user.Tasks[index+1:]...)
			break
		}
	}

}
