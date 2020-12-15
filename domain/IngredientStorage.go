package domain

import (
	"errors"
	"fmt"
)

type IngredientInventory struct {
	itemsWithQuantity map[Ingredient]Quantity
}

func NewIngredientInventory(items map[Ingredient]Quantity) *IngredientInventory {
	return &IngredientInventory{itemsWithQuantity: items}
}

func (inventory IngredientInventory) checkAvailability(ingredientReq map[Ingredient]Quantity) error {
	fmt.Printf("Available %+v\n", inventory.itemsWithQuantity)
	fmt.Printf("Required %+v\n", ingredientReq)

	for ingredient, _ := range ingredientReq {
		if _, ok := inventory.itemsWithQuantity[ingredient]; !ok {
			return errors.New(fmt.Sprintf("%s is not available", string(ingredient)))
		}
	}
	for ingredient, reqQuantity := range ingredientReq {
		if inventory.itemsWithQuantity[ingredient] == 0 {
			return errors.New(fmt.Sprintf("%s is not available", string(ingredient)))
		} else if inventory.itemsWithQuantity[ingredient] < reqQuantity {
			return errors.New(fmt.Sprintf("item %s is not sufficient", string(ingredient)))
		}
	}
	return nil
}

func (inventory IngredientInventory) update(ingredientReq map[Ingredient]Quantity) {
	for ingredient, reqQuantity := range ingredientReq {
		availableQuantity := inventory.itemsWithQuantity[ingredient]
		inventory.itemsWithQuantity[ingredient] = availableQuantity - reqQuantity
	}
}
