package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
	"fmt"
)

func FillUsers() {
	for i := 1; i <= 10; i++ {
		pkg.Users = append(pkg.Users, models.User{
			Name:        fmt.Sprintf("User%d", i),
			BirthYear:   1990 + i,
			PhoneNumber: fmt.Sprintf("9876543%d", i),
			Mail:        fmt.Sprintf("user%d@gmail.com", i),
			Address:     fmt.Sprintf("Some Avenue, %d", i),
			Balance:     10000,
			City:        pkg.Cities[i%len(pkg.Cities)],
		})
	}
}
