package service

import "go.uber.org/dig"

type (
	MoneyTransferSvc interface{}

	MoneyTransferSvcImpl struct {
		dig.In
	}
)

func NewMoneyTransferSvc(impl MoneyTransferSvcImpl) MoneyTransferSvc {
	return &impl
}
