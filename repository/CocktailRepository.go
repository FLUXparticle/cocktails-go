package repository

import "gorm.io/gorm"

// CocktailRepository handles database operations for cocktails.
type CocktailRepository struct {
	db *gorm.DB
}

// NewCocktailRepository creates a new instance of CocktailRepository.
func NewCocktailRepository(db *gorm.DB) *CocktailRepository {
	return &CocktailRepository{db: db}
}

// FindAll retrieves all cocktails from the database.
func (r *CocktailRepository) FindAll() []*Cocktail {
	var cocktails []*Cocktail
	if err := r.db.Find(&cocktails).Error; err != nil {
		panic(err)
	}
	return cocktails
}

// FindByID retrieves a cocktail by ID from the database.
func (r *CocktailRepository) FindByID(cocktailID uint) *Cocktail {
	var cocktail Cocktail
	if err := r.db.Preload("Instructions.Ingredient").First(&cocktail, cocktailID).Error; err != nil {
		panic(err)
	}
	return &cocktail
}

// FindByIDs retrieves multiple cocktails by their IDs from the database.
func (r *CocktailRepository) FindByIDs(cocktailIDs []uint) []Cocktail {
	var cocktails []Cocktail
	if err := r.db.Find(&cocktails, cocktailIDs).Error; err != nil {
		panic(err)
	}
	return cocktails
}

// FindByNameContains retrieves cocktails whose names contain the specified query.
func (r *CocktailRepository) FindByNameContains(query string) []*Cocktail {
	var cocktails []*Cocktail
	if err := r.db.Where("name LIKE ?", "%"+query+"%").Find(&cocktails).Error; err != nil {
		panic(err)
	}
	return cocktails
}

func (r *CocktailRepository) FindByIngredientIds(ingredientIDs []uint) []*Cocktail {
	var cocktails []*Cocktail
	err := r.db.
		Joins("JOIN instructions ON cocktails.cocktail_id = instructions.cocktail_id").
		Where("instructions.ingredient_id IN ?", ingredientIDs).
		Distinct().
		Find(&cocktails).Error
	if err != nil {
		panic(err)
	}
	return cocktails
}
