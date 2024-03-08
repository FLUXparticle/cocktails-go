package handler

import (
	"cocktails-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CocktailHandler struct {
	service *service.CocktailService
}

func NewCocktailHandler(service *service.CocktailService) *CocktailHandler {
	return &CocktailHandler{service: service}
}

func (h *CocktailHandler) CocktailList(c *gin.Context) {
	cocktails := h.service.GetAllCocktails()
	c.JSON(200, cocktails)
}

func (h *CocktailHandler) CocktailDetails(c *gin.Context) {
	idStr := c.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		cocktail := h.service.GetCocktailInstructions(uint(id))
		c.JSON(http.StatusOK, cocktail)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *CocktailHandler) IngredientList(c *gin.Context) {
	ingredients := h.service.GetAllIngredients()
	c.JSON(200, ingredients)
}

func (h *CocktailHandler) IngredientDetails(c *gin.Context) {
	idStr := c.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		ingredients, cocktails := h.service.GetCocktailsWithIngredients([]uint{uint(id)})
		c.JSON(http.StatusOK, gin.H{"name": ingredients[0].Name, "cocktails": cocktails})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *CocktailHandler) SearchCocktails(c *gin.Context) {
	query := c.Query("query")

	cocktails := h.service.Search(query)

	c.JSON(200, cocktails)
}
