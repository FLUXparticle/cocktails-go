package repository

type Cocktail struct {
	CocktailID   uint           `gorm:"primaryKey" json:"id"`
	Name         string         `json:"name"`
	Instructions []*Instruction `gorm:"foreignKey:CocktailID" json:"instructions,omitempty"`
}
