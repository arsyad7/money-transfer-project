package postgres

import "go.uber.org/dig"

type (
	MoneyTransferRepo interface{}

	MoneyTransferRepoImpl struct {
		dig.In
	}
)

func NewMoneyTransferRepo(impl MoneyTransferRepoImpl) MoneyTransferRepo {
	return &impl
}
