package entity

type User struct {
	Id       int
	Username string
	Password string
	Salt     string
	Email    string
}

type UserSlice []*User
