package domain

import "fmt"

type Outlet struct {
}

func (o *Outlet) dispenseBeverage(beverage Beverage) {
	fmt.Printf("%s is Prepared\n", string(beverage))
}
