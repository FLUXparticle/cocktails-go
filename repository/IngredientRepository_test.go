package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIngredientRepository(t *testing.T) {
	RunTest(t, func(repository *IngredientRepository) {
		// TODO repository ist nit nil
		assert.Fail(t, "not implemented")
	})
}

func TestIngredientRepository_FindAll(t *testing.T) {
	RunTest(t, func(repository *IngredientRepository) {
		// TODO Es gibt genau 61 Zutaten
		assert.Fail(t, "not implemented")
	})
}

func TestIngredientRepository_FindByNameContains(t *testing.T) {
	RunTest(t, func(repository *IngredientRepository) {
		// TODO Es gibt genau 24 Zutaten deren Namen "sirup" enthält
		assert.Fail(t, "not implemented")
	})
}
