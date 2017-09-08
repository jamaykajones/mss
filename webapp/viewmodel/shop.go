package viewmodel

import (
	"fmt"

	"github.com/jamaykajones/mss/webapp/model"
)

type Shop struct { //vm
	Title      string
	Active     string
	Categories []Category
}

type Category struct { //vm
	URL           string
	ImageURL      string
	Title         string
	Description   string
	IsOrientRight bool
}

func NewShop(categories []model.Category) Shop { // constructor func
	result := Shop{ // returns shop vm
		Title:  "Meme Shirt Supply - Shop",
		Active: "shop",
	}
	result.Categories = make([]Category, len(categories))
	for i := 0; i < len(categories); i++ {
		vm := categorytoVM(categories[i])
		vm.IsOrientRight = i%2 == 1
		result.Categories[i] = vm
	}
	return result
}

func categorytoVM(c model.Category) Category {
	return Category{
		URL:         fmt.Sprintf("/shop/%v", c.ID),
		ImageURL:    c.ImageURL,
		Title:       c.Title,
		Description: c.Description,
	}
}
