# Money Transfer Project Documentation

## Overview

This documentation provides information about the Money Transfer Project, including setup instructions, environment variables, and details about the mock APIs used.

## Project Setup

1. **Clone the repository:**
   ```bash
   git clone https://github.com/arsyad7/money-transfer-project
   ```
2. **Navigate to the project directory:**
    ```bash
    cd money-transfer-project
    ```
3. **Create table database:**
    
    You can copy the query from `table.sql` and execute on your local.
4. **run command:**
    ```bash
    go run .\cmd\money-transfer-project\main.go
    ```

## Environment Variables
Make sure to set the following environment variable before running the project, you can copy ENV file from `.env.example`

And in this case, for the `MOCKAPI_BASE_URL` you can use this value `https://65d3685d522627d50108d550.mockapi.io/v1`

## MockAPIs
The Money Transfer Project utilizes the following mock API endpoints:
1. **Account Validation Endpoint:**
    - URL: https://65d3685d522627d50108d550.mockapi.io/v1/account-validation
    - Description: This endpoint is used for validating user accounts.
2. **Transaction Endpoint:**
    - URL: https://65d3685d522627d50108d550.mockapi.io/v1/transaction
    - Description: This endpoint is used for handling money transactions.

## Running The Project
Once the setup and environment variables are configured, you can run the project using the following command:
```bash
go run .\cmd\money-transfer-project\main.go
```

## Code Flow
1. **GET localhost:8090/account-validation?accountName={accountName}&accountNumber={accountNumber}**
    Current data list
    ```bash
    [
        {
            "createdAt": "2024-02-19T05:15:54.511Z",
            "accountName": "Kristi Christiansen",
            "accountNumber": "82565507",
            "bankName": "Money Market Account",
            "id": "1"
        },
        {
            "createdAt": "2024-02-19T08:01:56.708Z",
            "accountName": "Gerald Treutel",
            "accountNumber": "93030766",
            "bankName": "Auto Loan Account",
            "id": "2"
        }
    ]
    ```
    Except data from the list will return error.

2. **POST localhost:8090/process-transaction**

    Request Body : 
    ```bash
    {
        "accountNumber": "93030766",
        "amount": -5000
    }
    ```
    Response Success : 
    ```bash
    {
        "TransactionID": "25",
        "Amount": -5000,
        "Status": "pending" // (initial value)
    }
    ```
    If you put `-` in amount, it will deduct user balance, if not, user balance will increase.
    You can get `accountNumber` from `/account-validation`

3. **POST localhost:8090/post-transaction**
    
    This endpoint is for callback endpoint from thirdparty.

    Request Body : 
    ```bash
    {
        "transactionID": "25",
        "amount": -5000,
        "status": "success"
    }
    ```

    Status value will be `success` or `failed`. If `success` account balance will increase and if `failed` there is no logic process. 
