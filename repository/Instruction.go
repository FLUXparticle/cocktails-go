package repository

type Instruction struct {
	CocktailID   uint        `json:"-"`
	AmountCL     int         `json:"amount"`
	IngredientID uint        `json:"-"`
	Ingredient   *Ingredient `gorm:"foreignKey:IngredientID" json:"ingredient"`
}
