package main

import (
	"Udemy/pkg"
	"Udemy/pkg/mock"
	"fmt"
)

func main() {
	mock.FillCategories()
	mock.FillCourses()
	mock.FillRegions()
	mock.FillCities()
	mock.FillUsers()
	mock.FillSubscriptions()
	for {
		var choice int

		fmt.Println("1. Получить курсы пользователя")
		fmt.Println("2. Получить общий рассход за курсы")
		fmt.Println("3. Поиск курсов по категории")
		fmt.Println("4. Поиск пользователей по региону")
		fmt.Println("5. Взять средний возраст пользователей по курсу")
		fmt.Println("6. Выйти")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			pkg.GetUserCourses()
		case 2:
			pkg.GetTotalSpentByUser()
		case 3:
			pkg.GetCoursesByCategory()
		case 4:
			pkg.GetUsersByRegion()
		case 5:
			pkg.GetAverageAgeByCourse()
		case 6:
			return
		}
	}
}
