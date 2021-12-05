package model

import "fmt"

type User struct {
	Name string
	Id   string
}

func NewUser(id, name string) *User {
	return &User{
		Name: name,
		Id:   id,
	}
}

func (u *User) GetId() string {
	return u.Id
}

func (u *User) ReceiveNotification(message string) {
	fmt.Println(fmt.Sprintf("%v received notification: %v", u.Name, message))
}
