package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CocktailHandler struct {
}

func NewCocktailHandler() *CocktailHandler {
	return &CocktailHandler{}
}

func (h *CocktailHandler) CocktailList(c *gin.Context) {
	panic("implement me")
}

func (h *CocktailHandler) CocktailDetails(c *gin.Context) {
	idStr := c.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		fmt.Println("id:", id)
		panic("implement me")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *CocktailHandler) IngredientList(c *gin.Context) {
	panic("implement me")
}

func (h *CocktailHandler) IngredientDetails(c *gin.Context) {
	idStr := c.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		fmt.Println("id:", id)
		panic("implement me")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (h *CocktailHandler) SearchCocktails(c *gin.Context) {
	query := c.Query("query")
	fmt.Println("query:", query)
	panic("implement me")
}
