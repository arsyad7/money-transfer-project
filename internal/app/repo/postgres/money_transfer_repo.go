package postgres

import (
	"context"
	"database/sql"

	"go.uber.org/dig"
)

type (
	UserBank struct {
		ID             int
		AccountName    string
		AccountNumber  string
		BankCode       string
		AccountBalance float64
	}

	Transaction struct {
		UserBankID      int
		TransactionID   string
		Amount          float64
		Status          string
		TransactionType string
	}

	MoneyTransferRepo interface {
		CreateUserBank(ctx context.Context, userBank UserBank) (err error)
		GetUserBankByAccountNumber(ctx context.Context, accountNumber string) (userBank UserBank, err error)
		CreateTransaction(ctx context.Context, transaction Transaction) (err error)
		UpdateUserBankBalance(ctx context.Context, userBank UserBank) (err error)
		GetTransactionByID(ctx context.Context, transactionID string) (transaction Transaction, err error)
		UpdateStatusTransaction(ctx context.Context, transactionID string, status string) (err error)
		GetUserBankByID(ctx context.Context, id int) (userBank UserBank, err error)
	}

	MoneyTransferRepoImpl struct {
		dig.In

		*sql.DB
	}
)

func NewMoneyTransferRepo(impl MoneyTransferRepoImpl) MoneyTransferRepo {
	return &impl
}

func (r *MoneyTransferRepoImpl) CreateUserBank(ctx context.Context, userBank UserBank) (err error) {
	_, err = r.ExecContext(ctx, `
	INSERT INTO money_transfer.user_banks (account_name, account_number, bank_code, account_balance)
	SELECT $1::text, $2::text, $3::text, $4::numeric
	WHERE NOT EXISTS (
		SELECT 1 FROM money_transfer.user_banks WHERE account_number = $2
	)`,
		userBank.AccountName, userBank.AccountNumber, userBank.BankCode, userBank.AccountBalance)
	if err != nil {
		return err
	}

	return nil
}

func (r *MoneyTransferRepoImpl) GetUserBankByAccountNumber(ctx context.Context, accountNumber string) (userBank UserBank, err error) {
	err = r.QueryRowContext(ctx, "SELECT id, account_name, account_number, bank_code, account_balance FROM money_transfer.user_banks WHERE account_number = $1", accountNumber).Scan(&userBank.ID, &userBank.AccountName, &userBank.AccountNumber, &userBank.BankCode, &userBank.AccountBalance)
	if err != nil {
		return userBank, err
	}

	return userBank, nil
}

func (r *MoneyTransferRepoImpl) CreateTransaction(ctx context.Context, transaction Transaction) (err error) {
	_, err = r.ExecContext(ctx, "INSERT INTO money_transfer.transactions (user_bank_id, transaction_id, amount, status, transaction_type) VALUES ($1, $2, $3, $4, $5)", transaction.UserBankID, transaction.TransactionID, transaction.Amount, transaction.Status, transaction.TransactionType)
	if err != nil {
		return err
	}

	return nil
}

func (r *MoneyTransferRepoImpl) UpdateUserBankBalance(ctx context.Context, userBank UserBank) (err error) {
	_, err = r.ExecContext(ctx, "UPDATE money_transfer.user_banks SET account_balance = account_balance + $1 WHERE id = $2", userBank.AccountBalance, userBank.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *MoneyTransferRepoImpl) GetTransactionByID(ctx context.Context, transactionID string) (transaction Transaction, err error) {
	err = r.QueryRowContext(ctx, "SELECT user_bank_id, transaction_id, amount, status, transaction_type FROM money_transfer.transactions WHERE transaction_id = $1", transactionID).Scan(&transaction.UserBankID, &transaction.TransactionID, &transaction.Amount, &transaction.Status, &transaction.TransactionType)
	if err != nil {
		if err == sql.ErrNoRows {
			return transaction, nil
		}

		return transaction, err
	}

	return transaction, nil
}

func (r *MoneyTransferRepoImpl) UpdateStatusTransaction(ctx context.Context, transactionID string, status string) (err error) {
	_, err = r.ExecContext(ctx, "UPDATE money_transfer.transactions SET status = $1 WHERE transaction_id = $2", status, transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (r *MoneyTransferRepoImpl) GetUserBankByID(ctx context.Context, id int) (userBank UserBank, err error) {
	err = r.QueryRowContext(ctx, "SELECT id, account_name, account_number, bank_code, account_balance FROM money_transfer.user_banks WHERE id = $1", id).Scan(&userBank.ID, &userBank.AccountName, &userBank.AccountNumber, &userBank.BankCode, &userBank.AccountBalance)
	if err != nil {
		return userBank, err
	}

	return userBank, nil
}
