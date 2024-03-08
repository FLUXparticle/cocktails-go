package repository

type Ingredient struct {
	IngredientID uint   `gorm:"primaryKey" json:"id"`
	Name         string `json:"name"`
}
