package models

type User struct {
	ID         int    `json:"id" db:"ID"`
	FirstName  string `json:"first_name" db:"FirstName"`
	LastName   string `json:"last_name" db:"LastName"`
	Patronymic string `json:"patronymic" db:"Patronymic"`
	SNILS      string `json:"snils" db:"SNILS"`
}
