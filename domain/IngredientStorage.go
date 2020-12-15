package domain

type IngredientInventory struct {
	itemsWithQuantity map[Ingredient]Quantity
}

func (is IngredientInventory) checkAvailability(ingredientReq map[Ingredient]Quantity) error {
	return nil
}

func (is IngredientInventory) update(quantity map[Ingredient]Quantity) {

}
