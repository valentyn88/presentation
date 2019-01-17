package mock

import (
	"github.com/valentyn88/presentation/domain"
)

// SimpleSearchSvc - implementation of ProductSearcher for test.
type SimpleSearchSvc struct {
}

// Search - search products.
func (s SimpleSearchSvc) Search(qp domain.QueryParam) ([]domain.Product, int64, error) {
	pp := products()

	return pp, int64(len(pp)), nil
}
