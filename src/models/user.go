package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Model
	Firstname string `json:"first_name"`
	Username  string `json:"username" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	//Tasks     []Task `json:"tasks" gorm:"foreignKey:TaskId"` //user will contain his unique slice of tasks
	//no need for this, I'd rather store UserId in tasks
}

/**
* Encrypts and hashes user's password.
*
* params password - unencrypted password
*
 */
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

/**
* Compares encrypted password with the input password
*
* params - unencrypted password
* Returns nil on success, or an error on failure.
 */
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
