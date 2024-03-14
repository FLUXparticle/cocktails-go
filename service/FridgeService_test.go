package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFridgeService(t *testing.T) {
	RunTest(t, func(service *FridgeService) {
		assert.NotNil(t, service)
	})
}

func TestFridgeService_CRUD(t *testing.T) {
	RunTest(t, func(service *FridgeService) {
		fridgeIngredients0 := service.GetFridgeIngredients()
		assert.Equal(t, 0, countInFridge(fridgeIngredients0))

		service.UpdateFridgeIngredient(29, true)

		fridgeIngredients1 := service.GetFridgeIngredients()
		assert.Equal(t, 1, countInFridge(fridgeIngredients1))

		service.UpdateFridgeIngredient(29, false)

		fridgeIngredients2 := service.GetFridgeIngredients()
		assert.Equal(t, 0, countInFridge(fridgeIngredients2))
	})
}

func TestFridgeService_GetPossibleCocktails(t *testing.T) {
	RunTest(t, func(service *FridgeService) {
		possibleCocktailsBefore := service.GetPossibleCocktails()
		assert.Len(t, possibleCocktailsBefore, 0)

		service.UpdateFridgeIngredient(8, true)
		service.UpdateFridgeIngredient(31, true)

		possibleCocktailsAfter := service.GetPossibleCocktails()
		assert.Len(t, possibleCocktailsAfter, 1)
	})
}

func TestFridgeService_GetShoppingList(t *testing.T) {
	RunTest(t, func(service *FridgeService) {
		ingredientsBefore := service.GetShoppingList()
		assert.Len(t, ingredientsBefore, 61)

		service.UpdateFridgeIngredient(29, true)

		ingredientsAfter := service.GetShoppingList()
		assert.Len(t, ingredientsAfter, 60)
	})
}

func countInFridge(fridgeIngredients []*FridgeIngredient) int {
	cnt := 0
	for _, ingredient := range fridgeIngredients {
		if ingredient.InFridge {
			cnt++
		}
	}
	return cnt
}
