package model

import "fmt"

//insert shop_deatils.html here
type Product struct {
	Name        string
	Material    string
	Description string
	Size        string
	Price       float32
	IsOrganic   bool
	ImageURL    string
	ID          int
	CategoryID  int
}

func GetProductsForCategory(categoryID int) []Product {
	result := []Product{}
	for _, p := range products {
		if p.CategoryID == categoryID {
			result = append(result, p)
		}
	}
	return result
}

func GetProduct(productID int) (*Product, error) {
	for _, p := range products {
		if p.ID == productID {
			return &p, nil
		}
	}
	return nil, fmt.Errorf("Product not found with ID %v", productID)
}

var products []Product = []Product{Product{
	Name:        "Awesome Shirt",
	Description: "When your awesome and you know it!",
	Material:    "Cotton. Grey.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "awsm.jpg",
	ID:          1,
	CategoryID:  1,
}, Product{
	Name:        "Fat Shirt",
	Description: "A tiny body couldn't store all this personality",
	Material:    "Cotton. Grey.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "fat.jpg",
	ID:          2,
	CategoryID:  1,
}, Product{
	Name:        "Hashtag Shirt",
	Description: "Blessed",
	Material:    "Cotton. Black.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "hstg.jpg",
	ID:          3,
	CategoryID:  1,
}, Product{
	Name:        "Ki_ _ Shirt",
	Description: "Fill in the blank, options my vary",
	Material:    "Cotton. White.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "ki_.jpg",
	ID:          4,
	CategoryID:  1,
}, Product{
	Name:        "Tupacca Shirt",
	Description: "Thug life or may the force be with you",
	Material:    "Cotton. White.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "pacca.jpg",
	ID:          5,
	CategoryID:  1,
}, Product{
	Name:        "Please Shirt",
	Description: "Please stop, you're killing me",
	Material:    "Cotton. Blue.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "pls.jpg",
	ID:          6,
	CategoryID:  1,
}, Product{
	Name:        "Shrug Shirt",
	Description: "Classic shrug emoji",
	Material:    "Cotton. White.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "shrg.jpg",
	ID:          7,
	CategoryID:  1,
}, Product{
	Name:        "Star Wars Shirt",
	Description: "Ones memes strong are",
	Material:    "Cotton. Black.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "sw.jpg",
	ID:          8,
	CategoryID:  1,
}, Product{
	Name:        "Taco Shirt",
	Description: "Yea I'm into fitness",
	Material:    "Cotton. Black.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "taco.jpg",
	ID:          9,
	CategoryID:  1,
}, Product{
	Name:        "Time Shirt",
	Description: "No time whatsoever",
	Material:    "Cotton. Tan.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "time.jpg",
	ID:          10,
	CategoryID:  1,
}, Product{
	Name:        "Wifi Shirt",
	Description: "Home is where wifi connects automatically",
	Material:    "Cotton. Black / White.",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "wf.jpg",
	ID:          11,
	CategoryID:  1,
}, Product{
	Name:        "Challenge Shirt",
	Description: "It's on now!",
	Material:    "Cotton. White",
	Size:        "S / M / L / XL",
	Price:       19.99,
	IsOrganic:   true,
	ImageURL:    "e.e.jpg",
	ID:          12,
	CategoryID:  1,
},
}
