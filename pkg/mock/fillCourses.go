package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillCourses() {
	pkg.Course = append(pkg.Course, models.Course{
		Name:     "Golang",
		Price:    900,
		Duration: 3,
		Category: []models.Category{pkg.Category[0]},
	})
	pkg.Course = append(pkg.Course, models.Course{
		Name:     "Python",
		Price:    750,
		Duration: 3,
		Category: []models.Category{pkg.Category[0]},
	})
	pkg.Course = append(pkg.Course, models.Course{
		Name:     "DigitalMarketing",
		Price:    500,
		Duration: 2,
		Category: []models.Category{
			pkg.Category[3],
			pkg.Category[2],
		},
	})
	pkg.Course = append(pkg.Course, models.Course{
		Name:     "AdobePhotoshop",
		Price:    600,
		Duration: 2,
		Category: []models.Category{pkg.Category[1]},
	})
	pkg.Course = append(pkg.Course, models.Course{
		Name:     "AdobeIllustrator",
		Price:    600,
		Duration: 2,
		Category: []models.Category{pkg.Category[1]},
	})
	pkg.Course = append(pkg.Course, models.Course{
		Name:     "TheCompleteDigitalMarketingCourse",
		Price:    1000,
		Duration: 4,
		Category: []models.Category{
			pkg.Category[3]},
	})
	pkg.Course = append(pkg.Course, models.Course{
		Name:     "TheBasicOfBusiness",
		Price:    500,
		Duration: 2,
		Category: []models.Category{
			pkg.Category[2],
		},
	})
}
