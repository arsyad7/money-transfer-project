package controller

import (
	"encoding/json"
	"money-transfer-project/internal/app/repo/rest"
	"money-transfer-project/internal/app/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	MoneyTransferCtrl interface {
		GetAccountValidation(ec echo.Context) (err error)
		ProcessTransaction(ec echo.Context) (err error)
		PostTransaction(ec echo.Context) (err error)
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

func (c *MoneyTransferCtrlImpl) ProcessTransaction(ec echo.Context) (err error) {
	request := rest.TransferMoneyRequest{}

	err = json.NewDecoder(ec.Request().Body).Decode(&request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx := ec.Request().Context()

	resp, err := c.MoneyTransferSvc.ProcessTransaction(ctx, request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, resp)
}

func (c *MoneyTransferCtrlImpl) PostTransaction(ec echo.Context) (err error) {
	request := service.PostTransactionRequest{}

	err = json.NewDecoder(ec.Request().Body).Decode(&request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx := ec.Request().Context()

	err = c.MoneyTransferSvc.PostTransaction(ctx, request)
	if err != nil {
		return ec.JSON(http.StatusInternalServerError, err.Error())
	}

	return ec.JSON(http.StatusOK, "Transaction has been posted")
}
