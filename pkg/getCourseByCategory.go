package pkg

import "fmt"

func GetCoursesByCategory() {
	var categoryName string
	fmt.Println("Введите имя категории")
	fmt.Scan(&categoryName)

	for _, course := range Course {
		for _, category := range course.Category {
			if category.Name == categoryName {
				fmt.Printf("Курс: %v\nЦена: %vc\n", course.Name, course.Price)
			}
		}
	}
}
