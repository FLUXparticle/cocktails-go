package repository

type Cocktail struct {
	CocktailID   uint `gorm:"primaryKey"`
	Name         string
	Instructions []*Instruction `gorm:"foreignKey:CocktailID"`
}
