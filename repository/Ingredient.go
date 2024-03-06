package repository

type Ingredient struct {
	IngredientID uint `gorm:"primaryKey"`
	Name         string
}
