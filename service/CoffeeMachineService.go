package service

import (
	. "CoffeeMachineDunzo/domain"
	"fmt"
)

type CoffeeMachineService struct {
	coffeeMachine *CoffeeMachine
}

func NewCoffeeMachineService(machine *CoffeeMachine) *CoffeeMachineService {
	return &CoffeeMachineService{
		coffeeMachine: machine,
	}
}

func (service *CoffeeMachineService) MakeBeverage(beverage Beverage, selectedOutlet int) ResponseMessage {
	err := service.coffeeMachine.MakeBeverage(beverage, selectedOutlet)
	if err != nil {
		return ResponseMessage(fmt.Sprintf("%s cannot be prepared because %s", beverage, err))
	}
	return ResponseMessage(fmt.Sprintf("%s is prepared", string(beverage)))
}

type ResponseMessage string
