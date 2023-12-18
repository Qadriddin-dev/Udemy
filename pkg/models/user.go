package models

type User struct {
	Name        string
	BirthYear   int
	PhoneNumber string
	Mail        string
	Address     string
	Balance     int
	City        City
}
