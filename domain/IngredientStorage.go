package domain

import (
	"errors"
	"fmt"
)

type IngredientInventory struct {
	itemsWithQuantity map[Ingredient]Quantity
	Publisher
}

func NewIngredientInventory(items map[Ingredient]Quantity, publisher Publisher) *IngredientInventory {
	return &IngredientInventory{itemsWithQuantity: items, Publisher: publisher}
}

func (inventory IngredientInventory) checkAvailability(ingredientReq map[Ingredient]Quantity) error {
	for ingredient, _ := range ingredientReq {
		if _, ok := inventory.itemsWithQuantity[ingredient]; !ok {
			return errors.New(fmt.Sprintf("%s is not available", string(ingredient)))
		}
	}
	for ingredient, reqQuantity := range ingredientReq {
		if inventory.itemsWithQuantity[ingredient] == 0 {
			return errors.New(fmt.Sprintf("%s is not available", string(ingredient)))
		} else if inventory.itemsWithQuantity[ingredient] < reqQuantity {
			inventory.NotifyAll(ingredient)
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

func (inventory IngredientInventory) NotifyAll(ingredient Ingredient) {
	for _, o := range inventory.subscribers {
		o.notify(ingredient)
	}
}
