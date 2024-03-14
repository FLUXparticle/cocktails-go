package main

import (
	"cocktails-go/handler"
	"cocktails-go/repository"
	"cocktails-go/service"
	"context"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
	"moul.io/zapgorm2"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func StaticFileServer(c *gin.Context) {
	path := c.Request.URL.Path
	// Überprüfe, ob der Pfad auf ".html" oder ".js" endet
	if path == "/" || strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".js") {
		// Versuche, die Datei aus dem "web"-Ordner zu servieren
		filePath := "./web" + path
		if _, err := os.Stat(filePath); err == nil {
			http.ServeFile(c.Writer, c.Request, filePath)
			c.Abort() // Stoppe die Weiterleitung der Anfrage
		}
	}
}

func NewGinHandler(cocktailsHandler *handler.CocktailHandler, fridgeHandler *handler.FridgeHandler, log *zap.Logger) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(ginzap.Ginzap(log, time.RFC3339, true))
	r.Use(gin.Recovery())

	r.Use(StaticFileServer)

	// CocktailsHandler-Endpunkte
	r.GET("/api/cocktails", cocktailsHandler.CocktailList)
	r.GET("/api/cocktails/:id", cocktailsHandler.CocktailDetails)
	r.GET("/api/ingredients", cocktailsHandler.IngredientList)
	r.GET("/api/ingredients/:id", cocktailsHandler.IngredientDetails)
	r.GET("/api/search", cocktailsHandler.SearchCocktails)

	// FridgeHandler-Endpunkte
	r.GET("/api/user/fridge", fridgeHandler.FridgeList)
	r.PATCH("/api/user/fridge/:ingredientId", fridgeHandler.UpdateFridgeIngredient)
	r.GET("/api/user/possible", fridgeHandler.PossibleRecipes)
	r.GET("/api/user/shopping", fridgeHandler.ShoppingList)

	return r
}

func NewHTTPServer(lc fx.Lifecycle, handler http.Handler, log *zap.Logger) *http.Server {
	srv := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: handler,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			//fmt.Println("Starting HTTP server at", srv.Addr)
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func NewZapLogger() *zap.Logger {
	//zapLogger, _ := zap.NewProduction(zap.AddCaller() /*, zap.AddCallerSkip(1)*/)
	//return zapLogger
	return zap.NewExample(zap.AddCaller() /*, zap.AddCallerSkip(1)*/)
}

func NewFxLogger(log *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: log}
}

func NewDatabaseLogger(log *zap.Logger) logger.Interface {
	return zapgorm2.New(log).LogMode(logger.Info)
}

func main() {
	fx.New(
		//fx.WithLogger(NewFxLogger),
		fx.Provide(
			NewHTTPServer,
			NewGinHandler,
			NewZapLogger,
			NewDatabaseLogger,
			handler.NewCocktailHandler,
			handler.NewFridgeHandler,
			service.NewCocktailService,
			service.NewFridgeService,
			repository.NewCocktailRepository,
			repository.NewIngredientRepository,
			repository.NewDatabase,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
