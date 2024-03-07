package service

import "cocktails-go/repository"

// CocktailService repräsentiert den Service für Cocktails und deren Funktionen.
type CocktailService struct {
	cocktailRepo   *repository.CocktailRepository
	ingredientRepo *repository.IngredientRepository
}

// NewCocktailService erstellt eine neue Instanz des CocktailService mit den erforderlichen Repositories.
func NewCocktailService(cocktailRepo *repository.CocktailRepository, ingredientRepo *repository.IngredientRepository) *CocktailService {
	return &CocktailService{
		cocktailRepo:   cocktailRepo,
		ingredientRepo: ingredientRepo,
	}
}

// GetAllCocktails gibt eine Liste aller Cocktails zurück.
func (s *CocktailService) GetAllCocktails() []*repository.Cocktail {
	return s.cocktailRepo.FindAll()
}

// GetAllIngredients gibt eine Liste aller Zutaten zurück.
func (s *CocktailService) GetAllIngredients() []*repository.Ingredient {
	return s.ingredientRepo.FindAll()
}

// GetCocktailInstructions gibt eine Liste der Anweisungen für einen bestimmten Cocktail zurück.
func (s *CocktailService) GetCocktailInstructions(cocktailID uint) *repository.Cocktail {
	return s.cocktailRepo.FindByID(cocktailID)
}

// GetCocktailsWithIngredients gibt eine Liste der Zutaten zurück und aller Cocktails, die jeweils mindestens eine der angegebenen Zutaten enthalten.
func (s *CocktailService) GetCocktailsWithIngredients(ingredientIDs []uint) ([]*repository.Ingredient, []*repository.Cocktail) {
	ingredients := s.ingredientRepo.FindByIDs(ingredientIDs)
	cocktails := s.cocktailRepo.FindByIngredientIds(ingredientIDs)
	return ingredients, cocktails
}

// Search führt eine Suche nach Cocktails durch, die den Query entweder im Namen oder im Namen einer der Zutaten enthalten.
func (s *CocktailService) Search(query string) []*repository.Cocktail {
	cocktails1 := s.cocktailRepo.FindByNameContains(query)
	ingredients := s.ingredientRepo.FindByNameContains(query)

	ingredientIDs := make([]uint, len(ingredients))
	for i, ingredient := range ingredients {
		ingredientIDs[i] = ingredient.IngredientID
	}

	cocktails2 := s.cocktailRepo.FindByIngredientIds(ingredientIDs)

	// Erstelle eine Map, um Duplikate zu vermeiden
	cocktailMap := make(map[uint]*repository.Cocktail)

	// Füge Cocktails aus cocktails1 zur Map hinzu
	for _, cocktail := range cocktails1 {
		cocktailMap[cocktail.CocktailID] = cocktail
	}

	// Füge Cocktails aus cocktails2 zur Map hinzu, überschreibe dabei ggf. vorhandene Einträge
	for _, cocktail := range cocktails2 {
		cocktailMap[cocktail.CocktailID] = cocktail
	}

	// Konvertiere die Map zurück in ein Slice
	cocktails := make([]*repository.Cocktail, 0, len(cocktailMap))
	for _, cocktail := range cocktailMap {
		cocktails = append(cocktails, cocktail)
	}

	return cocktails
}
