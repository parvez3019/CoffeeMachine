package domain

type Ingredient string

type Quantity int

type Recipe struct {
	ingredientsWithQuantity map[Ingredient]Quantity
}

func NewRecipe(ingredientsWithQuantity map[Ingredient]Quantity) *Recipe {
	return &Recipe{ingredientsWithQuantity: ingredientsWithQuantity}
}
