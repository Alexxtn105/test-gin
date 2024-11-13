package models

type User struct {
	// TODO
	ID   uint64
	Name string
}

func UserFind(userId uint64) *User {
	//TODO
	return &User{}
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
