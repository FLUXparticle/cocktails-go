package handler

import (
	"cocktails-go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FridgeHandler struct {
}

func NewFridgeHandler() *FridgeHandler {
	return &FridgeHandler{}
}

func (h *FridgeHandler) FridgeList(c *gin.Context) {
	panic("implement me")
}

func (h *FridgeHandler) UpdateFridgeIngredient(c *gin.Context) {
	var err error
	idStr := c.Param("ingredientId")
	var id int
	if id, err = strconv.Atoi(idStr); err == nil {
		fmt.Println("id:", id)
		var fridgeIngredient service.FridgeIngredient
		if err = c.ShouldBindJSON(&fridgeIngredient); err == nil {
			panic("implement me")
			c.JSON(http.StatusOK, nil)
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func (h *FridgeHandler) PossibleRecipes(c *gin.Context) {
	panic("implement me")
}

func (h *FridgeHandler) ShoppingList(c *gin.Context) {
	panic("implement me")
}
