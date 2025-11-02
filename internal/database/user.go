package database


type User struct {
	ID int
	FirstName string
	LastName string
	Phone string
}


func NewUser(id int, firstName string, lastName string, phone string) *User {
	return &User{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		Phone: phone,
	}
}
