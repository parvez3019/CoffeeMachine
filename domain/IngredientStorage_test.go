package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldReturnNoErrorIfIngredientsAreAvailable(t *testing.T) {
	ingredientInventory := setupInventory()
	requestedIngredients := map[Ingredient]Quantity{
		"hot_water":        500,
		"hot_milk":         500,
		"ginger_syrup":     100,
		"sugar_syrup":      100,
		"tea_leaves_syrup": 100,
	}
	assert.Nil(t, ingredientInventory.checkAvailability(requestedIngredients))
}

func Test_ShouldReturnErrorIfIngredientsAreInsufficient(t *testing.T) {
	ingredientInventory := setupInventory()
	requestedIngredients := map[Ingredient]Quantity{
		"hot_water":        600,
		"hot_milk":         500,
	}
	assert.EqualError(t, ingredientInventory.checkAvailability(requestedIngredients), "item hot_water is not sufficient")
}

func Test_ShouldReturnErrorIfIngredientsAreNotAvailable(t *testing.T) {
	ingredientInventory := setupInventory()
	requestedIngredients := map[Ingredient]Quantity{
		"chocolate":        600,
	}
	assert.EqualError(t, ingredientInventory.checkAvailability(requestedIngredients), "chocolate is not available")
}

func setupInventory() *IngredientInventory {
	inventories := map[Ingredient]Quantity{
		"hot_water":        500,
		"hot_milk":         500,
		"ginger_syrup":     100,
		"sugar_syrup":      100,
		"tea_leaves_syrup": 100,
	}

	publisher := Publisher{}
	publisher.AddSubscriber(NewAlertSubscriber())
	return NewIngredientInventory(inventories, publisher)
}
