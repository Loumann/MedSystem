package models

import "sync"

type User struct {
	ID         int    `json:"id" db:"ID"`
	FirstName  string `json:"first_name" db:"FirstName"`
	LastName   string `json:"last_name" db:"LastName"`
	Patronymic string `json:"patronymic" db:"Patronymic"`
	SNILS      string `json:"snils" db:"SNILS"`
}

type WaitingUsers struct {
	sync.RWMutex
	Items []User
}

type FulfillWaitingUserInput struct {
	User    int      `json:"user" binding:"required"`
	Analyse Analysis `json:"analyse" binding:"required"`
}
