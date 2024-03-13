package repository

import (
	"cocktails-go/test"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"testing"
)

func RunTest(t *testing.T, block any) {
	// Fx-App für Tests starten
	app := fxtest.New(t,
		fx.Provide(
			test.NewTestLogger,
			NewDatabase,
			NewIngredientRepository,
			NewCocktailRepository,
		),
		fx.Invoke(block),
	)
	defer app.RequireStop()
	app.RequireStart()
}
