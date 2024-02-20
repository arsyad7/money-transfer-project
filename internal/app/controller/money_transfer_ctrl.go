package controller

import (
	"money-transfer-project/internal/app/repo/rest"
	"money-transfer-project/internal/app/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	MoneyTransferCtrl interface {
		GetAccountValidation(ec echo.Context) (err error)
	}

	MoneyTransferCtrlImpl struct {
		dig.In

		MoneyTransferSvc service.MoneyTransferSvc
	}
)

func NewMoneyTransferCtrl(impl MoneyTransferCtrlImpl) MoneyTransferCtrl {
	return &impl
}

func (c *MoneyTransferCtrlImpl) GetAccountValidation(ec echo.Context) (err error) {
	request := rest.GetAccountValidationRequest{}

	request.AccountName = ec.FormValue("accountName")
	request.AccountNumber = ec.FormValue("accountNumber")

	ctx := ec.Request().Context()

	resp, err := c.MoneyTransferSvc.GetAccountValidation(ctx, request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, resp)
}
