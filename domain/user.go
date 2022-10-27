package domain

type User struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Name     string `db:"name"`
}
