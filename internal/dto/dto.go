package dto

type CreateProductInput struct {
	Name  string  `jason:"name"`
	Price float64 `jason:"price"`
}
