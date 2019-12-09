package db

import "github.com/kodran/workshop/day3/webapi/model"

// Client table model
var Client = []model.Client{
	model.Client{
		ID:       1,
		Name:     "Jorge",
		LastName: "Castro",
		Age:      30,
	},
	model.Client{
		ID:       2,
		Name:     "Roberto",
		LastName: "Garcia",
		Age:      30,
	},
	model.Client{
		ID:       3,
		Name:     "Max",
		LastName: "Deboli",
		Age:      36,
	},
}
