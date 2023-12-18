package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillRegions() {
	pkg.Region = append(pkg.Region, models.Region{
		Name: "Dushanbe",
	})
	pkg.Region = append(pkg.Region, models.Region{
		Name: "Sughd",
	})
	pkg.Region = append(pkg.Region, models.Region{
		Name: "Khatlon",
	})
}
