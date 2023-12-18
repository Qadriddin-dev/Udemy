package pkg

import "fmt"

func GetUsersByRegion() {
	var regionName string
	fmt.Println("Введите название региона")
	fmt.Scan(&regionName)

	for _, user := range Users {
		if user.City.Region.Name == regionName {
			fmt.Printf("Пользователь: %v\n", user.Name)
		}
	}
}
