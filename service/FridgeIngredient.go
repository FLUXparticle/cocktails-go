package service

import "cocktails-go/repository"

type FridgeIngredient struct {
	*repository.Ingredient
	InFridge bool
}
