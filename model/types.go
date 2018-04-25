package model

type UserCreatedEvent struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
