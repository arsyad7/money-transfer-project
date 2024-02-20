package app

import (
	"money-transfer-project/internal/app/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func setRoute(
	e *echo.Echo,

	moneyTransferCtrl controller.MoneyTransferCtrl,
) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/account-validation", moneyTransferCtrl.GetAccountValidation)
	e.POST("/process-transaction", moneyTransferCtrl.ProcessTransaction)
	e.POST("/post-transaction", moneyTransferCtrl.PostTransaction)
}
