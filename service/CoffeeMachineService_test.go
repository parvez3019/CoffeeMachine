package service

import (
	. "CoffeeMachineDunzo/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldMakeBeverage(t *testing.T) {
	coffeeMachineService := NewCoffeeMachineService(createCoffeeMachine())

	assert.Equal(t, ResponseMessage("hot_tea is prepared"),
		coffeeMachineService.MakeBeverage("hot_tea", 0))

	assert.Equal(t, ResponseMessage("hot_coffee is prepared"),
		coffeeMachineService.MakeBeverage("hot_coffee", 1))

	assert.Equal(t, ResponseMessage("green_tea cannot be prepared because green_mixture is not available"),
		coffeeMachineService.MakeBeverage("green_tea", 2))

	assert.Equal(t, ResponseMessage("black_tea cannot be prepared because item hot_water is not sufficient"),
		coffeeMachineService.MakeBeverage("black_tea", 3))
}

func createCoffeeMachine() *CoffeeMachine {
	builder := NewCoffeeMachineBuilder()
	return builder.
		AddInventory(setupInventory()).
		AddRecipes(setupRecipes()).
		SetTotalOutlets(setupOutlets()).
		Build()
}

func setupOutlets() []Outlet {
	outlets := make([]Outlet, 4)
	for i, _ := range outlets {
		outlets[i] = Outlet{}
	}
	return outlets
}

func setupRecipes() *Recipes {
	recipes := map[Beverage]*Recipe{
		"hot_tea": NewRecipe(map[Ingredient]Quantity{
			"hot_water":        200,
			"hot_milk":         100,
			"ginger_syrup":     10,
			"sugar_syrup":      10,
			"tea_leaves_syrup": 30,
		}),
		"hot_coffee": NewRecipe(map[Ingredient]Quantity{
			"hot_water":        100,
			"ginger_syrup":     30,
			"hot_milk":         400,
			"sugar_syrup":      50,
			"tea_leaves_syrup": 30,
		}),
		"black_tea": NewRecipe(map[Ingredient]Quantity{
			"hot_water":        300,
			"ginger_syrup":     30,
			"sugar_syrup":      50,
			"tea_leaves_syrup": 30,
		}),
		"green_tea": NewRecipe(map[Ingredient]Quantity{
			"hot_water":     100,
			"ginger_syrup":  30,
			"sugar_syrup":   50,
			"green_mixture": 30,
		}),
	}
	return CreateRecipes(recipes)
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
