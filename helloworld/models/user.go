package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextId = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.Id != 0 {
		return User{}, errors.New("User must not contain id")
	}
	u.Id = nextId
	nextId++
	users = append(users, &u)
	return u, nil
}

func GetUserById(id int) (User, error) {
	for _, v := range users {
		if v.Id == id {
			return *v, nil
		}
	}
	return User{}, fmt.Errorf("User with id '%v' not found", id)
}

func UpdateUser(id int, u User) (User, error) {
	for i, v := range users {
		if v.Id == id {
			users[i] = &u
			users[i].Id = id
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with id '%v' not found", u.Id)
}

func DeleteUserById(id int) error {
	for i, v := range users {
		if v.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id '%v' not found", id)
}
