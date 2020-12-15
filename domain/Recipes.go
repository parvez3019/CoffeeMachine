package domain

import "errors"

type Recipes struct {
	recipes map[Beverage]*Recipe
}

func (r *Recipes) get(beverage Beverage) (*Recipe, error) {
	recipe := r.recipes[beverage]
	if recipe == nil {
		return nil, errors.New("Invalid_Beverage")
	}
	return recipe, nil
}
