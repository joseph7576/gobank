package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

// constructor
func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(100),
		FirstName: firstName,
		LastName:  lastName,
		Number:    rand.Int63n(100000000),
	}
}
