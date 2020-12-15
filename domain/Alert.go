package domain

import (
	"fmt"
)

type Subscriber interface {
	notify(Ingredient)
}

type AlertService struct {
}

func NewAlertSubscriber() Subscriber {
	return &AlertService{}
}

func (as *AlertService) notify(ingredient Ingredient) {
	fmt.Printf("Ingredient %s running low, please refill\n", ingredient)
}
