package service

import (
	"cocktails-go/repository"
	"cocktails-go/test"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func RunTest(t *testing.T, block any) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal(r)
		}
	}()

	// Fx-App für Tests starten
	app := fxtest.New(t,
		fx.Provide(
			test.NewTestLogger,
			repository.NewDatabase,
			repository.NewIngredientRepository,
			repository.NewCocktailRepository,
			NewCocktailService,
			NewFridgeService,
		),
		fx.Invoke(block),
	)

	defer app.RequireStop()
	app.RequireStart()
}
