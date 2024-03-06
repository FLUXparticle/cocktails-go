package main

import (
	"bufio"
	"cocktails-go/repository"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strconv"
	"strings"
)

func readCocktailDataset(txtFilename string) []*repository.Cocktail {
	var cocktails []*repository.Cocktail

	ingredientsMap := make(map[string]*repository.Ingredient)

	in, err := os.Open(txtFilename)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	var cocktail *repository.Cocktail
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		line := sc.Text()
		if cocktail == nil {
			cocktail = &repository.Cocktail{Name: line}
		} else if len(line) > 0 {
			var amountCL int
			var ingredientName string
			split := strings.Split(line, "cl:")
			switch len(split) {
			case 1:
				ingredientName = split[0]
			case 2:
				if amountCL, err = strconv.Atoi(split[0]); err != nil {
					panic(err)
				} else {
					ingredientName = split[1]
				}
			}

			ingredient, found := ingredientsMap[ingredientName]
			if !found {
				ingredientID := uint(len(ingredientsMap) + 1)
				ingredient = &repository.Ingredient{
					IngredientID: ingredientID,
					Name:         ingredientName,
				}
				ingredientsMap[ingredientName] = ingredient
			}

			cocktail.Instructions = append(cocktail.Instructions, &repository.Instruction{
				AmountCL:   amountCL,
				Ingredient: ingredient,
			})
		} else {
			cocktails = append(cocktails, cocktail)
			cocktail = nil
		}
	}

	return cocktails
}

func writeCocktailDataset(dbFilename string, cocktails []*repository.Cocktail) {
	os.Remove(dbFilename)

	cfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	db, err := gorm.Open(sqlite.Open(dbFilename), cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("--> db.AutoMigrate(&repository.Instruction{})")
	db.AutoMigrate(&repository.Instruction{})

	fmt.Println("--> db.AutoMigrate(&repository.Cocktail{})")
	db.AutoMigrate(&repository.Cocktail{})

	fmt.Println("--> db.Create(ds.cocktails)")
	db.Create(cocktails)
}

func main() {
	cocktails := readCocktailDataset("cocktails.txt")
	writeCocktailDataset("cocktails.db", cocktails)
}
