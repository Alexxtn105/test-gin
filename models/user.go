package models

type User struct {
	ID   int
	Name string
}

func UserCheck(email string, password string) *User {

	return &User{}
}
func UserCheckAvailability(email string) bool {

	return true
}
