package repository

import "gorm.io/gorm"

// IngredientRepository handles database operations for ingredients.
type IngredientRepository struct {
	db *gorm.DB
}

// NewIngredientRepository creates a new instance of IngredientRepository.
func NewIngredientRepository(db *gorm.DB) *IngredientRepository {
	return &IngredientRepository{db: db}
}

// FindAll retrieves all ingredients from the database.
func (r *IngredientRepository) FindAll() []*Ingredient {
	var ingredients []*Ingredient
	if err := r.db.Find(&ingredients).Error; err != nil {
		panic(err)
	}
	return ingredients
}

func (r *IngredientRepository) FindByIDs(ingredientIDs []uint) []*Ingredient {
	var ingredients []*Ingredient
	if err := r.db.Find(&ingredients, ingredientIDs).Error; err != nil {
		panic(err)
	}
	return ingredients
}

// FindByNameContains retrieves ingredients whose names contain the specified query.
func (r *IngredientRepository) FindByNameContains(query string) []Ingredient {
	var ingredients []Ingredient
	if err := r.db.Where("name LIKE ?", "%"+query+"%").Find(&ingredients).Error; err != nil {
		panic(err)
	}
	return ingredients
}
