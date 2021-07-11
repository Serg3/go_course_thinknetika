package users

type User struct {
	id       int
	name     string
	password string
	admin    bool
}

type Users struct {
	list []User
}

func New() *Users {
	u := Users{}
	u.list = append(u.list,
		User{id: 1, name: "root", password: "root", admin: true},
		User{id: 2, name: "usr", password: "123456", admin: false},
	)
	return &u
}

func (usr *User) ID() int {
	return usr.id
}

func (usr *User) Admin() bool {
	return usr.admin
}

func (u *Users) Search(name, psw string) *User {
	for _, usr := range u.list {
		if usr.name == name && usr.password == psw {
			return &usr
		}
	}
	return nil
}
