package mock

import (
	"Udemy/pkg"
	"Udemy/pkg/models"
)

func FillSubscriptions() {
	for i := range pkg.Users {
		if i < len(pkg.Course) {
			pkg.Users[i].Balance -= pkg.Course[i].Price

			subscription := models.Subscription{
				User:   pkg.Users[i],
				Course: pkg.Course[i],
				Price:  pkg.Course[i].Price,
			}
			pkg.Subscriptions = append(pkg.Subscriptions, subscription)
		}
	}
}
