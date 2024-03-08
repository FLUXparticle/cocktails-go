package handler

import (
	"cocktails-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FridgeHandler struct {
	service *service.FridgeService
}

func NewFridgeHandler(service *service.FridgeService) *FridgeHandler {
	return &FridgeHandler{service: service}
}

func (h *FridgeHandler) FridgeList(c *gin.Context) {
	fridgeIngredients := h.service.GetFridgeIngredients()
	c.JSON(200, fridgeIngredients)
}

func (h *FridgeHandler) UpdateFridgeIngredient(c *gin.Context) {
	var err error
	idStr := c.Param("ingredientId")
	var id int
	if id, err = strconv.Atoi(idStr); err == nil {
		var fridgeIngredient service.FridgeIngredient
		if err = c.ShouldBindJSON(&fridgeIngredient); err == nil {
			h.service.UpdateFridgeIngredient(uint(id), fridgeIngredient.InFridge)
			c.JSON(http.StatusOK, nil)
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func (h *FridgeHandler) PossibleRecipes(c *gin.Context) {
	cocktails := h.service.GetPossibleCocktails()
	c.JSON(200, cocktails)
}

func (h *FridgeHandler) ShoppingList(c *gin.Context) {
	shoppingList := h.service.GetShoppingList()
	c.JSON(200, shoppingList)
}
