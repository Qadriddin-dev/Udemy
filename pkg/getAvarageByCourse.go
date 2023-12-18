package pkg

import (
	"fmt"
)

func GetAverageAgeByCourse() {
	var courseName string
	fmt.Println("Введите название курса")
	fmt.Scan(&courseName)

	totalAge := 0
	count := 0
	for _, subscription := range Subscriptions {
		if subscription.Course.Name == courseName {
			age := 2023 - subscription.User.BirthYear
			totalAge += age
			count++
		}
	}

	if count > 0 {
		averageAge := totalAge / count
		fmt.Printf("Средний возраст: %v\n", averageAge)
	} else {
		fmt.Println("Нет подписок на указанный курс.")
	}
}
