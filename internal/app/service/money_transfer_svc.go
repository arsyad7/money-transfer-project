package service

import (
	"context"
	"fmt"
	"money-transfer-project/internal/app/repo/postgres"
	"money-transfer-project/internal/app/repo/rest"

	"go.uber.org/dig"
)

type (
	ProcessTransactionResponse struct {
		TransactionID string
		Amount        float64
		Status        string
	}

	PostTransactionRequest struct {
		TransactionID string  `json:"transactionID"`
		Amount        float64 `json:"amount"`
		Status        string  `json:"status"`
	}

	MoneyTransferSvc interface {
		GetAccountValidation(ctx context.Context, req rest.GetAccountValidationRequest) (resp []rest.GetAccountValidationResponse, err error)
		ProcessTransaction(ctx context.Context, req rest.TransferMoneyRequest) (resp ProcessTransactionResponse, err error)
		PostTransaction(ctx context.Context, req PostTransactionRequest) (err error)
	}

	MoneyTransferSvcImpl struct {
		dig.In

		MockAPIRepo       rest.MockAPIRepo
		MoneyTransferRepo postgres.MoneyTransferRepo
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

	for _, v := range resp {
		err = s.MoneyTransferRepo.CreateUserBank(ctx, postgres.UserBank{
			AccountName:    v.AccountName,
			AccountNumber:  v.AccountNumber,
			BankCode:       v.BankName,
			AccountBalance: 0,
		})
		if err != nil {
			return resp, err
		}
	}

	return resp, nil
}

func (s *MoneyTransferSvcImpl) ProcessTransaction(ctx context.Context, req rest.TransferMoneyRequest) (resp ProcessTransactionResponse, err error) {
	if req.AccountNumber == "" {
		return resp, fmt.Errorf("account number is required")
	}

	if req.Amount == 0 {
		return resp, fmt.Errorf("cant process transaction with 0 amount")
	}

	if req.TransactionType == "" && req.Amount != 0 {
		if req.Amount < 0 {
			req.TransactionType = rest.TransactionTypeCredit
		} else {
			req.TransactionType = rest.TransactionTypeDebit
		}
	}

	userBank, err := s.MoneyTransferRepo.GetUserBankByAccountNumber(ctx, req.AccountNumber)
	if err != nil {
		return resp, err
	}

	transaction, err := s.MockAPIRepo.TransferMoney(ctx, req)
	if err != nil {
		return resp, err
	}

	transactionReq := postgres.Transaction{
		UserBankID:      userBank.ID,
		TransactionID:   transaction.TransactionID,
		Amount:          req.Amount,
		TransactionType: string(req.TransactionType),
		Status:          "pending",
	}
	err = s.MoneyTransferRepo.CreateTransaction(ctx, transactionReq)
	if err != nil {
		return resp, err
	}

	resp.Amount = req.Amount
	resp.TransactionID = transaction.TransactionID
	resp.Status = "pending"

	return resp, nil
}

func (s *MoneyTransferSvcImpl) PostTransaction(ctx context.Context, req PostTransactionRequest) (err error) {
	if req.TransactionID == "" {
		return fmt.Errorf("transaction id is required")
	}

	if req.Amount == 0 {
		return fmt.Errorf("cant process transaction with 0 amount")
	}

	transaction, err := s.MoneyTransferRepo.GetTransactionByID(ctx, req.TransactionID)
	if err != nil {
		return err
	}

	if transaction.TransactionID == "" {
		return fmt.Errorf("transaction not found")
	}

	if transaction.Status == "success" {
		return fmt.Errorf("transaction already success")
	}

	err = s.MoneyTransferRepo.UpdateStatusTransaction(ctx, req.TransactionID, req.Status)
	if err != nil {
		return err
	}

	if req.Status == "success" {
		userBank, err := s.MoneyTransferRepo.GetUserBankByID(ctx, transaction.UserBankID)
		if err != nil {
			return err
		}

		err = s.MoneyTransferRepo.UpdateUserBankBalance(ctx, postgres.UserBank{
			ID:             userBank.ID,
			AccountBalance: transaction.Amount,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
