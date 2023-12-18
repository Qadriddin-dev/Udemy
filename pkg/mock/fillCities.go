package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillCities() {
	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Dushanbe",
		Region: pkg.Region[0],
	})
	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Khujand",
		Region: pkg.Region[1],
	})
	pkg.Cities = append(pkg.Cities, models.City{
		Name:   "Kulob",
		Region: pkg.Region[2],
	})
}
