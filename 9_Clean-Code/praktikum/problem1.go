package main

type user struct {
	id       int
	username int
	password int
}
//
type userService struct {
	t []user
}

func (u userService) getAllUsers() []user {
	return u.t
}

func (u userService) getUserById(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}
	return user{}
}
