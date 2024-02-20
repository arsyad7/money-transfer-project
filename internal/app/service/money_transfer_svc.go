package service

import (
	"context"
	"money-transfer-project/internal/app/repo/rest"

	"go.uber.org/dig"
)

type (
	MoneyTransferSvc interface {
		GetAccountValidation(ctx context.Context, req rest.GetAccountValidationRequest) (resp []rest.GetAccountValidationResponse, err error)
	}

	MoneyTransferSvcImpl struct {
		dig.In

		MockAPIRepo rest.MockAPIRepo
	}
)

func NewMoneyTransferSvc(impl MoneyTransferSvcImpl) MoneyTransferSvc {
	return &impl
}

func (s *MoneyTransferSvcImpl) GetAccountValidation(ctx context.Context, req rest.GetAccountValidationRequest) (resp []rest.GetAccountValidationResponse, err error) {
	resp, err = s.MockAPIRepo.GetAccountValidation(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
