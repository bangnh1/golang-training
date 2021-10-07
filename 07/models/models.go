package models

import (
	fake "github.com/brianvoe/gofakeit/v6"
)

type User struct {
	Id       int    `fake:"{Id:0,30}"`
	FullName string `fake:"{name}"`
	Email    string
	Phone    string
	Age      int `fake:"{Age:1,100}"`
	Sex      string
}

var (
	users []*User
)

func init() {
	users = []*User{}

	for i := 0; i < 30; i++ {
		users = append(users, &User{
			Id:       int(fake.Number(0, 30)),
			FullName: fake.Name(),
			Email:    fake.Email(),
			Phone:    fake.Phone(),
			Age:      int(fake.Number(1, 100)),
			Sex:      fake.Gender(),
		})
	}
}

func ListUser() []*User {
	return users
}

type ByAge []*User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ById []*User

func (a ById) Len() int           { return len(a) }
func (a ById) Less(i, j int) bool { return a[i].Id < a[j].Id }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByName []*User

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].FullName < a[j].FullName }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
