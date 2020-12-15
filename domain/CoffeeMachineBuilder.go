package domain

type CoffeeMachineBuilder struct {
	*CoffeeMachine
}

func NewCoffeeMachineBuilder() *CoffeeMachineBuilder {
	return &CoffeeMachineBuilder{
		CoffeeMachine: NewCoffeeMachine(),
	}
}

func (builder *CoffeeMachineBuilder) addInventory(inventory *IngredientInventory) *CoffeeMachineBuilder {
	builder.ingredientInventory = inventory
	return builder
}

func (builder *CoffeeMachineBuilder) addRecipes(recipes Recipes) *CoffeeMachineBuilder {
	builder.recipes = recipes
	return builder
}

func (builder *CoffeeMachineBuilder) setTotalOutlets(outlets []Outlet) *CoffeeMachineBuilder {
	builder.outlets = outlets
	return builder
}

func (builder *CoffeeMachineBuilder) build() *CoffeeMachine {
	return builder.CoffeeMachine
}
