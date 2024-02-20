package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"go.uber.org/dig"
)

type TransactionType string

const (
	TransactionTypeDebit  TransactionType = "debit"
	TransactionTypeCredit TransactionType = "credit"
)

type (
	GetAccountValidationRequest struct {
		AccountNumber string `json:"accountNumber"`
		AccountName   string `json:"accountName"`
	}

	GetAccountValidationResponse struct {
		CreatedAt     string `json:"createdAt"`
		AccountName   string `json:"accountName"`
		AccountNumber string `json:"accountNumber"`
		BankName      string `json:"bankName"`
		ID            string `json:"id"`
	}

	TransferMoneyRequest struct {
		AccountNumber   string          `json:"accountNumber"`
		Amount          float64         `json:"amount"`
		TransactionType TransactionType `json:"transactionType"`
	}

	TransferMoneyResponse struct {
		Amount        float64 `json:"amount"`
		TransactionID string  `json:"id"`
	}

	MockAPIRepo interface {
		GetAccountValidation(ctx context.Context, req GetAccountValidationRequest) (resp []GetAccountValidationResponse, err error)
		TransferMoney(ctx context.Context, req TransferMoneyRequest) (resp TransferMoneyResponse, err error)
	}

	MockAPIRepoImpl struct {
		dig.In
	}
)

func NewMockAPIRepo(impl MockAPIRepoImpl) MockAPIRepo {
	return &impl
}

func (r *MockAPIRepoImpl) GetAccountValidation(ctx context.Context, req GetAccountValidationRequest) (resp []GetAccountValidationResponse, err error) {
	baseURL := fmt.Sprintf("%s/account-validation?", os.Getenv("MOCKAPI_BASE_URL"))

	urlA, err := url.Parse(baseURL)
	values := urlA.Query()

	if req.AccountNumber != "" {
		values.Add("accountNumber", req.AccountNumber)
	}

	if req.AccountName != "" {
		values.Add("accountName", req.AccountName)
	}

	urlA.RawQuery = values.Encode()

	res, err := http.Get(urlA.String())
	if err != nil {
		return resp, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return resp, fmt.Errorf("akun bank tidak ditemukan. Kode status: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return resp, err
	}

	return resp, nil
}

func (r *MockAPIRepoImpl) TransferMoney(ctx context.Context, req TransferMoneyRequest) (resp TransferMoneyResponse, err error) {
	baseURL := fmt.Sprintf("%s/transaction", os.Getenv("MOCKAPI_BASE_URL"))

	if req.Amount == 0 {
		return resp, fmt.Errorf("cant process transaction with 0 amount")
	}

	// Create the request body
	reqBody := map[string]interface{}{
		"amount": req.Amount,
	}

	// Convert the request body to JSON
	reqBodyJSON, err := json.Marshal(reqBody)
	if err != nil {
		return resp, err
	}

	// Perform the HTTP POST request
	res, err := http.Post(baseURL, "application/json", bytes.NewBuffer(reqBodyJSON))
	if err != nil {
		return resp, err
	}

	defer res.Body.Close()

	// Check the HTTP status code
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return resp, fmt.Errorf("transaction failed. Status code: %d", res.StatusCode)
	}

	// Decode the JSON response into the resp variable
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return resp, err
	}

	return resp, nil
}
