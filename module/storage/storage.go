package storage

import "github.com/valentyn88/presentation/module/product"

// Storager - interface for operations with Product.
type Storager interface {
	Search(qp product.QueryParam) ([]product.Product, int64, error)
}
