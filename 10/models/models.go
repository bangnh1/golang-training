package models

import (
	"fmt"
	"time"
)

type User struct {
	tableName struct{}  `pg:"test.public.users"`
	Id        string    `pg:"id,pk" json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Age       string    `json:"age"`
	Sex       string    `json:"sex"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Age      string `json:"age"`
	Sex      string `json:"sex"`
}

/*
In ra thông tin của user
*/
func (u *User) Print() {
	fmt.Println("Id : ", u.Id)
	fmt.Println("FullName : ", u.FullName)
	fmt.Println("Email : ", u.Email)
	fmt.Println("Phone : ", u.Phone)
	fmt.Println("Age : ", u.Age)
	fmt.Println("Sex : ", u.Sex)
	fmt.Println("CreatedAt : ", u.CreatedAt)
	fmt.Println("UpdatedAt : ", u.UpdatedAt)
}

type ByAge []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByName []User

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].FullName < a[j].FullName }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
