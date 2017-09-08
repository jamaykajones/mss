package viewmodel

import (
	"html/template"

	"github.com/jamaykajones/mss/webapp/model"
)

type ProductViewModel struct {
	Title   string
	Active  string
	Product Product
}

type Product struct {
	Name        string
	Material    string
	Description template.HTML
	Size        string
	Price       float32
	IsOrganic   bool
	ImageURL    string
	ID          int
}

func productToVM(product *model.Product) Product {
	return Product{
		Name:        product.Name,
		Material:    product.Material,
		Description: template.HTML(product.Description),
		Size:        product.Size,
		Price:       product.Price,
		IsOrganic:   product.IsOrganic,
		ImageURL:    product.ImageURL,
		ID:          product.ID,
	}
}

func NewProduct(product *model.Product) ProductViewModel {
	return ProductViewModel{
		Title:   "Meme Shirt Supply - Shop",
		Active:  "shop",
		Product: productToVM(product),
	}
}
