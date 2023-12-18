package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillCategories() {
	pkg.Category = append(pkg.Category, models.Category{
		Name: "Programming",
	})
	pkg.Category = append(pkg.Category, models.Category{
		Name: "Design",
	})
	pkg.Category = append(pkg.Category, models.Category{
		Name: "Business",
	})
	pkg.Category = append(pkg.Category, models.Category{
		Name: "Marketing",
	})
}
