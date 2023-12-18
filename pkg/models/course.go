package models

type Course struct {
	Name     string
	Price    int
	Duration int
	Category []Category
}
