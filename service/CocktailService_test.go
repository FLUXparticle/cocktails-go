package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCocktailService(t *testing.T) {
	RunTest(t, func(service *CocktailService) {
		assert.NotNil(t, service)
	})
}

func TestCocktailService_GetAllCocktails(t *testing.T) {
	RunTest(t, func(service *CocktailService) {
		cocktails := service.GetAllCocktails()
		if assert.Len(t, cocktails, 69) {
			assert.Len(t, cocktails[0].Instructions, 0)
		}
	})
}

func TestCocktailService_GetAllIngredients(t *testing.T) {
	RunTest(t, func(service *CocktailService) {
		ingredients := service.GetAllIngredients()
		assert.Len(t, ingredients, 61)
	})

}

func TestCocktailService_GetCocktailInstructions(t *testing.T) {
	RunTest(t, func(service *CocktailService) {
		cocktail := service.GetCocktailInstructions(24)
		if assert.NotNil(t, cocktail) {
			assert.Equal(t, "Milkshake", cocktail.Name)
			assert.Len(t, cocktail.Instructions, 4)
		}
	})
}

func TestCocktailService_GetCocktailsWithIngredients(t *testing.T) {
	RunTest(t, func(service *CocktailService) {
		_, cocktails := service.GetCocktailsWithIngredients([]uint{29})
		if assert.Len(t, cocktails, 6) {
			assert.Len(t, cocktails[0].Instructions, 0)
		}
	})
}

func TestCocktailService_Search(t *testing.T) {
	RunTest(t, func(service *CocktailService) {
		cocktails := service.Search("Milch")
		if assert.Len(t, cocktails, 7) {
			assert.Len(t, cocktails[0].Instructions, 0)
		}
	})
}
