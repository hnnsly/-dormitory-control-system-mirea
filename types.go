package main

import "math/rand"

type Account struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewAccount(login, password string) *Account {
	return &Account{
		ID:       rand.Intn(1000),
		Login:    login,
		Password: password,
	}
}
