package domain

import (
	"errors"
	"sync"
)

type Beverage string

type BeverageMaker struct {
	ingredientInventory *IngredientInventory
	recipes             Recipes
	outlets             []Outlet
}

var beverageMaker *BeverageMaker = nil
var mutex = &sync.Mutex{}

func GetBeverageMachineInstance(outlets int) *BeverageMaker {
	if beverageMaker == nil {
		beverageMaker = newBeverageMaker(outlets)
	}
	return beverageMaker
}

func newBeverageMaker(outlets int) *BeverageMaker {
	return &BeverageMaker{
		ingredientInventory: &IngredientInventory{},
		recipes:             Recipes{},
		outlets:             make([]Outlet, outlets),
	}
}

func (bm *BeverageMaker) makeBeverage(beverage Beverage, outletNumber int) (bool, error) {
	if outletNumber >= len(bm.outlets) {
		return false, errors.New("Invalid_Outlet")
	}
	r, err := bm.recipes.get(beverage)
	if err == nil {
		return false, err
	}

	mutex.Lock()

	err = bm.ingredientInventory.checkAvailability(r.ingredientsWithQuantity)
	if err != nil {
		return false, err
	}
	bm.ingredientInventory.update(r.ingredientsWithQuantity)

	bm.outlets[outletNumber].dispenseBeverage(beverage)

	mutex.Unlock()
	return true, nil
}
