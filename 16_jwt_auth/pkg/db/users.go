package users

// User's information structure.
type User struct {
	id       int
	name     string
	password string
	admin    bool
}

// Users is a collection to store users.
type Users struct {
	list []User
}

// New populates Users collection.
func New() *Users {
	u := Users{}
	u.list = append(u.list,
		User{id: 1, name: "root", password: "root", admin: true},
		User{id: 2, name: "usr", password: "123456", admin: false},
	)
	return &u
}

// ID returns user's id.
func (usr *User) ID() int {
	return usr.id
}

// Admin determines if the user is an administrator.
func (usr *User) Admin() bool {
	return usr.admin
}

// User returns a user's address by name and password.
func (u *Users) User(name, psw string) *User {
	for _, usr := range u.list {
		if usr.name == name && usr.password == psw {
			return &usr
		}
	}
	return nil
}
