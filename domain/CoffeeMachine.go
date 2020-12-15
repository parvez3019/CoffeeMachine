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
	return &CoffeeMachine{}
}

func (cm *CoffeeMachine) makeBeverage(beverage Beverage, outletNumber int) (bool, error) {
	if outletNumber >= len(cm.outlets) {
		return false, errors.New("Invalid_Outlet")
	}
	r, err := cm.recipes.get(beverage)
	if err == nil || r == nil {
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
