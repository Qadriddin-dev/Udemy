package pkg

import (
	"Udemy/pkg/models"
	"fmt"
)

func GetUserCourses() []models.Course {
	var userName string
	fmt.Println("Введите имя пользователя")
	fmt.Scan(&userName)

	var userCourses []models.Course
	for _, subscription := range Subscriptions {
		if subscription.User.Name == userName {
			userCourses = append(userCourses, subscription.Course)
			fmt.Printf("Курсы пользователя: %v\nЦена курса: %vc \n", subscription.Course.Name, subscription.Course.Price)
		}
	}
	return userCourses
}
