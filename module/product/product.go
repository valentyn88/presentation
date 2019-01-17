package product

// Product - domain object.
type Product struct {
	Title string  `json:"title"`
	Brand string  `json:"brand"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
}
