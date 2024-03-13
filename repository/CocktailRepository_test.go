package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCocktailRepository(t *testing.T) {
	RunTest(t, func(repository *CocktailRepository) {
		assert.NotNil(t, repository)
	})
}

func TestCocktailRepository_FindAll(t *testing.T) {
	RunTest(t, func(repository *CocktailRepository) {
		cocktails := repository.FindAll()
		assert.Len(t, cocktails, 69)
	})
}

func TestCocktailRepository_FindByID(t *testing.T) {
	RunTest(t, func(repository *CocktailRepository) {
		cocktail := repository.FindByID(24)
		assert.Equal(t, "Milkshake", cocktail.Name)
		assert.Len(t, cocktail.Instructions, 4)
	})
}

func TestCocktailRepository_FindByNameContains(t *testing.T) {
	RunTest(t, func(repository *CocktailRepository) {
		cocktails := repository.FindByNameContains("Milk")
		assert.Len(t, cocktails, 3)
	})
}
