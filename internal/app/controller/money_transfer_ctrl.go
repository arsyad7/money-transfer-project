package controller

import "go.uber.org/dig"

type (
	MoneyTransferCtrl interface{}

	MoneyTransferCtrlImpl struct {
		dig.In
	}
)

func NewMoneyTransferCtrl(impl MoneyTransferCtrlImpl) MoneyTransferCtrl {
	return &impl
}
