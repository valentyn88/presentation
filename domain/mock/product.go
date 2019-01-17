package mock

import (
	"github.com/valentyn88/presentation/domain"
)

func products() []domain.Product {
	products := []domain.Product{
		domain.Product{
			Title: "Adidas sneakers",
			Brand: "Adidas",
			Price: 29.99,
			Stock: 5,
		},
		domain.Product{
			Title: "Nike sneakers",
			Brand: "Nike",
			Price: 39.99,
			Stock: 10,
		},
		domain.Product{
			Title: "Asics sneakers",
			Brand: "Asics",
			Price: 19.99,
			Stock: 15,
		},
	}

	return products
}
