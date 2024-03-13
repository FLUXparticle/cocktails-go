package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCocktailRepository(t *testing.T) {
	RunTest(t, func(repository *CocktailRepository) {
		// TODO repository ist nicht nil
		assert.Fail(t, "not implemented")
	})
}

func TestCocktailRepository_FindAll(t *testing.T) {
	RunTest(t, func(repository *CocktailRepository) {
		// TODO Es gibt genau 69 Cocktails
		assert.Fail(t, "not implemented")
	})
}

func TestCocktailRepository_FindByID(t *testing.T) {
	RunTest(t, func(repository *CocktailRepository) {
		// TODO Der Cocktail mit ID 24 heißt "Milkshake" und hat genau 4 Zutaten
		assert.Fail(t, "not implemented")
	})
}

func TestCocktailRepository_FindByNameContains(t *testing.T) {
	RunTest(t, func(repository *CocktailRepository) {
		// TODO Es gibt genau 3 Cocktails, deren Namen "Milk" enthält
		assert.Fail(t, "not implemented")
	})
}
