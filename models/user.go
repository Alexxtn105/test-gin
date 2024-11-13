package models

type User struct {
	// TODO
	ID   int
	Name string
}

func UserCheck(email string, password string) *User {
	//TODO
	return &User{}
}
func UserCheckAvailability(email string) bool {
	//TODO
	return true
}

func UserCreate(email string, password string) *User {
	//TODO
	return &User{}
}
