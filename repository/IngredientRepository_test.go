package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIngredientRepository(t *testing.T) {
	RunTest(t, func(repository *IngredientRepository) {
		assert.NotNil(t, repository)
	})
}

func TestIngredientRepository_FindAll(t *testing.T) {
	RunTest(t, func(repository *IngredientRepository) {
		ingredients := repository.FindAll()
		assert.Len(t, ingredients, 61)
	})
}

func TestIngredientRepository_FindByNameContains(t *testing.T) {
	RunTest(t, func(repository *IngredientRepository) {
		ingredients := repository.FindByNameContains("sirup")
		assert.Len(t, ingredients, 24)
	})
}
