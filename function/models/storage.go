package models

// Storager - interface for operations with Product.
type Storager interface {
	Search(qp QueryParam) ([]Product, int64, error)
}
