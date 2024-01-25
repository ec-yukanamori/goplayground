package model

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var Users = map[int]*User{
	1: {ID: 1, Name: "John Doe"},
	2: {ID: 2, Name: "Jane Doe"},
}
