package domain

type CoffeeMachineBuilder struct {
	*CoffeeMachine
}

func NewCoffeeMachineBuilder() *CoffeeMachineBuilder {
	return &CoffeeMachineBuilder{
		CoffeeMachine: NewCoffeeMachine(),
	}
}

func (builder *CoffeeMachineBuilder) AddInventory(inventory *IngredientInventory) *CoffeeMachineBuilder {
	builder.ingredientInventory = inventory
	return builder
}

func (builder *CoffeeMachineBuilder) AddRecipes(recipes *Recipes) *CoffeeMachineBuilder {
	builder.recipes = recipes
	return builder
}

func (builder *CoffeeMachineBuilder) SetTotalOutlets(outlets []Outlet) *CoffeeMachineBuilder {
	builder.outlets = outlets
	return builder
}

func (builder *CoffeeMachineBuilder) Build() *CoffeeMachine {
	return builder.CoffeeMachine
}
