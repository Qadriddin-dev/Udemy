package pkg

import "fmt"

func GetTotalSpentByUser() {
	var userName string
	fmt.Println("Введите имя пользователя")
	fmt.Scan(&userName)
	total := 0
	for _, subscription := range Subscriptions {
		if subscription.User.Name == userName {
			total += subscription.Price
		}
	}
	fmt.Printf("Расход за курс: %vc \n", total)
}
