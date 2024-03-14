package service

import "cocktails-go/repository"

// CocktailService repräsentiert den Service für Cocktails und deren Funktionen.
type CocktailService struct {
}

// NewCocktailService erstellt eine neue Instanz des CocktailService mit den erforderlichen Repositories.
func NewCocktailService() *CocktailService {
	return &CocktailService{}
}

// GetAllCocktails gibt eine Liste aller Cocktails zurück.
func (s *CocktailService) GetAllCocktails() []*repository.Cocktail {
	panic("implement me")
}

// GetAllIngredients gibt eine Liste aller Zutaten zurück.
func (s *CocktailService) GetAllIngredients() []*repository.Ingredient {
	panic("implement me")
}

// GetCocktailInstructions gibt eine Liste der Anweisungen für einen bestimmten Cocktail zurück.
func (s *CocktailService) GetCocktailInstructions(cocktailID uint) *repository.Cocktail {
	panic("implement me")
}

// GetCocktailsWithIngredients gibt eine Liste der Zutaten zurück und aller Cocktails, die jeweils mindestens eine der angegebenen Zutaten enthalten.
func (s *CocktailService) GetCocktailsWithIngredients(ingredientIDs []uint) ([]*repository.Ingredient, []*repository.Cocktail) {
	panic("implement me")
}

// Search führt eine Suche nach Cocktails durch, die den Query entweder im Namen oder im Namen einer der Zutaten enthalten.
func (s *CocktailService) Search(query string) []*repository.Cocktail {
	panic("implement me")
}
