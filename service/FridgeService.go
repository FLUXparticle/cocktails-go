package service

import "cocktails-go/repository"

// FridgeService repräsentiert den Service für den Kühlschrank und seine Funktionen.
type FridgeService struct {
	fridge map[uint]bool
}

// NewFridgeService erstellt eine neue Instanz des FridgeService mit den erforderlichen Repositories und Services.
func NewFridgeService() *FridgeService {
	return &FridgeService{
		fridge: make(map[uint]bool),
	}
}

// GetFridgeIngredients gibt eine Liste aller Zutaten zurück, zusammen mit einem Flag, ob sie im Kühlschrank sind oder nicht.
func (s *FridgeService) GetFridgeIngredients() []*FridgeIngredient {
	panic("implement me")
}

// UpdateFridgeIngredient aktualisiert den Status einer Zutat im Kühlschrank.
func (s *FridgeService) UpdateFridgeIngredient(ingredientID uint, inFridge bool) {
	panic("implement me")
}

// GetPossibleCocktails gibt eine Liste von Cocktails zurück, die mit den im Kühlschrank vorhandenen Zutaten zubereitet werden können.
func (s *FridgeService) GetPossibleCocktails() []*repository.Cocktail {
	panic("implement me")
}

// GetShoppingList gibt eine Liste aller Zutaten zurück, die die nicht im Kühlschrank sind.
func (s *FridgeService) GetShoppingList() []*repository.Ingredient {
	panic("implement me")
}
