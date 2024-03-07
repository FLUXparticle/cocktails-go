package service

import "cocktails-go/repository"

// FridgeService repräsentiert den Service für den Kühlschrank und seine Funktionen.
type FridgeService struct {
	fridge          map[uint]bool
	cocktailService *CocktailService
	ingredientRepo  *repository.IngredientRepository
}

// NewFridgeService erstellt eine neue Instanz des FridgeService mit den erforderlichen Repositories und Services.
func NewFridgeService(cocktailsService *CocktailService, ingredientRepo *repository.IngredientRepository) *FridgeService {
	return &FridgeService{
		fridge:          make(map[uint]bool),
		ingredientRepo:  ingredientRepo,
		cocktailService: cocktailsService,
	}
}

// GetFridgeIngredients gibt eine Liste aller Zutaten zurück, zusammen mit einem Flag, ob sie im Kühlschrank sind oder nicht.
func (s *FridgeService) GetFridgeIngredients() []*FridgeIngredient {
	ingredients := s.ingredientRepo.FindAll()

	fridgeIngredients := make([]*FridgeIngredient, len(ingredients))
	for i, ingredient := range ingredients {
		fridgeIngredients[i] = &FridgeIngredient{
			Ingredient: ingredient,
			InFridge:   s.fridge[ingredient.IngredientID],
		}
	}

	return fridgeIngredients
}

// UpdateFridgeIngredient aktualisiert den Status einer Zutat im Kühlschrank.
func (s *FridgeService) UpdateFridgeIngredient(ingredientID uint, inFridge bool) {
	if inFridge {
		s.fridge[ingredientID] = true
	} else {
		delete(s.fridge, ingredientID)
	}
}

// GetPossibleCocktails gibt eine Liste von Cocktails zurück, die mit den im Kühlschrank vorhandenen Zutaten zubereitet werden können.
func (s *FridgeService) GetPossibleCocktails() []*repository.Cocktail {
	ingredientIDs := make([]uint, 0, len(s.fridge))
	for id := range s.fridge {
		ingredientIDs = append(ingredientIDs, id)
	}

	// Holen Sie alle Cocktails aus der Datenbank
	_, eligibleCocktails := s.cocktailService.GetCocktailsWithIngredients(ingredientIDs)

	// Filtere die Cocktails basierend auf dem Kühlschrank
	possibleCocktails := make([]*repository.Cocktail, 0)

	for _, cocktail := range eligibleCocktails {
		// Überprüfe, ob alle Zutaten des Cocktails im Kühlschrank sind
		allIngredientsInFridge := true

		cocktail = s.cocktailService.GetCocktailInstructions(cocktail.CocktailID)

		for _, instruction := range cocktail.Instructions {
			ingredientID := instruction.Ingredient.IngredientID

			// Überprüfe, ob die Zutat im Kühlschrank ist
			if !s.fridge[ingredientID] {
				allIngredientsInFridge = false
				break
			}
		}

		// Füge den Cocktail zur Liste der möglichen Cocktails hinzu, wenn alle Zutaten im Kühlschrank sind
		if allIngredientsInFridge {
			possibleCocktails = append(possibleCocktails, cocktail)
		}
	}

	return possibleCocktails
}

// GetShoppingList gibt eine Liste aller Zutaten zurück, die die nicht im Kühlschrank sind.
func (s *FridgeService) GetShoppingList() []*repository.Ingredient {
	ingredients := s.ingredientRepo.FindAll()

	shoppingList := make([]*repository.Ingredient, 0, len(ingredients))

	for _, ingredient := range ingredients {
		if !s.fridge[ingredient.IngredientID] {
			shoppingList = append(shoppingList, ingredient)
		}
	}

	return shoppingList
}
