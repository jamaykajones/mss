package viewmodel

type Shop struct { //vm
	Title      string
	Active     string
	Categories []Category
}

type Category struct { //vm
	URL         string
	ImageURL    string
	Title       string
	Description string
}

func NewShop() Shop { // constructor func
	result := Shop{ // returns shop vm
		Title:  "Meme Shirt Supply - Shop",
		Active: "shop",
	}
	shirtCategory := Category{
		URL:      "/shop_details",
		ImageURL: "logo.png",
		Title:    "Shop Shirts",
		Description: `Explore our wide assortment of juices and mixes expected by
						today's lemonade stand clientelle. Now featuring a full line of
						organic juices that are guaranteed to be obtained from trees that
						have never been treated with pesticides or artificial
						fertilizers.`,
	}
	adCategory := Category{
		URL:      ".",
		ImageURL: "sign.png",
		Title:    "Shop Signs & Advertising",
		Description: `From paper cups to bio-degradable plastic to straws and
						napkins, LSS is your source for the sundries that keep your stand
						running smoothly.`,
	}
	createCategory := Category{
		URL:      ".",
		ImageURL: "create.png",
		Title:    "Create Your Own",
		Description: `Sure, you could just wait for people to find your stand
						along the side of the road, but if you want to take it to the next
						level, our premium line of advertising supplies.`,
	}
	result.Categories = []Category{shirtCategory, adCategory, createCategory}
	return result
}
