package domain

import (
	"errors"
	"sync"
)

type Beverage string

type CoffeeMachine struct {
	ingredientInventory *IngredientInventory
	recipes             Recipes
	outlets             []Outlet
}

var mutex = &sync.Mutex{}

func NewCoffeeMachine() *CoffeeMachine {
	return &CoffeeMachine{
		ingredientInventory: &IngredientInventory{},
		recipes:             Recipes{},
		outlets:             []Outlet{},
	}
}

func (cm *CoffeeMachine) addInventory(inventory *IngredientInventory) *CoffeeMachine {
	cm.ingredientInventory = inventory
	return cm
}

func (cm *CoffeeMachine) addRecipes(recipes Recipes) *CoffeeMachine {
	cm.recipes = recipes
	return cm
}

func (cm *CoffeeMachine) setTotalOutlets(outlets []Outlet) *CoffeeMachine {
	cm.outlets = outlets
	return cm
}

func (cm *CoffeeMachine) makeBeverage(beverage Beverage, outletNumber int) (bool, error) {
	if outletNumber >= len(cm.outlets) {
		return false, errors.New("Invalid_Outlet")
	}
	r, err := cm.recipes.get(beverage)
	if err == nil {
		return false, err
	}

	mutex.Lock()

	err = cm.ingredientInventory.checkAvailability(r.ingredientsWithQuantity)
	if err != nil {
		return false, err
	}
	cm.ingredientInventory.update(r.ingredientsWithQuantity)

	cm.outlets[outletNumber].dispenseBeverage(beverage)

	mutex.Unlock()
	return true, nil
}
