package domain

import (
	"errors"
	"fmt"
	"sync"
)

type Beverage string

type CoffeeMachine struct {
	ingredientInventory *IngredientInventory
	recipes             *Recipes
	outlets             []Outlet
}

var mutex = &sync.Mutex{}

var coffeeMachineInstance *CoffeeMachine = nil

func NewCoffeeMachine() *CoffeeMachine {
	if coffeeMachineInstance != nil {
		return coffeeMachineInstance
	}
	return &CoffeeMachine{
		ingredientInventory: &IngredientInventory{},
		recipes:             &Recipes{},
		outlets:             []Outlet{},
	}
}

func (cm *CoffeeMachine) MakeBeverage(beverage Beverage, outletNumber int) error {
	if outletNumber >= len(cm.outlets) {
		return errors.New("Invalid_Outlet")
	}
	r, err := cm.recipes.get(beverage)
	if err != nil {
		return err
	}

	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println(beverage)
	err = cm.ingredientInventory.checkAvailability(r.ingredientsWithQuantity)
	if err != nil {

		return err
	}
	cm.ingredientInventory.update(r.ingredientsWithQuantity)

	cm.outlets[outletNumber].dispenseBeverage(beverage)

	return nil
}
