package domain

type User struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Name     string `db:"name"`
}

func NewUser(username, email, password, name string) *User {

	return &User{
		Username: username,
		Email:    email,
		Password: password,
		Name:     name,
	}

}
