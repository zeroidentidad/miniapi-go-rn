package models

import (
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users []User

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{ID: 1, Username: "Zero2", Password: "xd2"}

	users[user.ID] = user
}

func GetUser(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return User{}, errors.New("Error: usuario no existe.")
}

func GetUsers() Users {
	list := Users{}

	for _, user := range users {
		list = append(list, user)
	}

	return list
}

func SaveUser(user User) User {
	user.ID = len(users) + 1
	users[user.ID] = user
	return user
}

func UpdateUser(user User, username, password string) User {
	user.Username = username
	user.Password = password
	users[user.ID] = user
	return user
}

func DeleteUser(id int) {
	delete(users, id)
}
