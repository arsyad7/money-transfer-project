package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"go.uber.org/dig"
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

	MockAPIRepo interface {
		GetAccountValidation(ctx context.Context, req GetAccountValidationRequest) (resp []GetAccountValidationResponse, err error)
	}

	MockAPIRepoImpl struct {
		dig.In
	}
)

func NewMockAPIRepo(impl MockAPIRepoImpl) MockAPIRepo {
	return &impl
}

func (r *MockAPIRepoImpl) GetAccountValidation(ctx context.Context, req GetAccountValidationRequest) (resp []GetAccountValidationResponse, err error) {
	base_url := fmt.Sprintf("%s/account-validation?", os.Getenv("MOCKAPI_BASE_URL"))

	urlA, err := url.Parse(base_url)
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
		return resp, fmt.Errorf("permintaan HTTP tidak berhasil. Kode status: %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return resp, err
	}

	return resp, nil
}
