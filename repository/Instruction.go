package repository

type Instruction struct {
	CocktailID   uint
	AmountCL     int
	IngredientID uint
	Ingredient   *Ingredient `gorm:"foreignKey:IngredientID"`
}
