// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import "github.com/kamva/mgm/v3"

type AuthPayload struct {
	User  *UserView `json:"user"`
	Token string    `json:"token"`
}

type LogInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID   string    `json:"id"`
	Text string    `json:"text"`
	Done bool      `json:"done"`
	User *UserView `json:"user"`
}

type User struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	Password         string `json:"password"`
}

type UserInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserView struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
