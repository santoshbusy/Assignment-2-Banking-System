# Banking System Simulation

## Overview
This project simulates a comprehensive banking system with support for multiple banks, branches, customers, savings accounts, and loans. It provides a RESTful API to manage these entities and perform financial transactions such as deposits, withdrawals, transfers, and loan repayments.

## Features

### bank & Branch Management
- **Create Banks**: Register new banking institutions.
- **Create Branches**: Add branches to specific banks.
- **View Details**: Retrieve details of banks and branches.

### Customer Management
- **Create Customers**: Register new customers with name and email.
- **View Profile**: Retrieve customer details including their accounts and loans.

### Account Operations
- **Open Account**: Create savings accounts for customers in specific branches.
- **Deposit & Withdraw**: Perform cash deposits and withdrawals with validation.
- **Fund Transfer**: Transfer money between accounts transactionally.
- **Transaction History**: View a history of all transactions for an account.

### Loan Management
- **Apply for Loan**: Customers can take loans from a bank branch.
- **Interest Calculation**: Loans have a fixed interest rate of 12%. View pending loan details and interest calculated for the current year.
- **Repay Loan**: Repay the loan amount (partial or full repayment).


## API Documentation

### General
- `GET /health`: Check system health status.

### Banks
- `POST /banks`: Create a new bank.
- `GET /banks`: List all banks.

### Branches
- `POST /branches`: Create a new branch.
- `GET /branches`: List all branches.

### Customers
- `POST /customers`: Create a new customer.
- `GET /customers/:id`: Get customer details.

### Accounts
- `POST /accounts`: Open a new savings account.
- `GET /accounts/:id`: Get account details.
- `POST /accounts/:id/deposit`: Deposit money into an account.
- `POST /accounts/:id/withdraw`: Withdraw money from an account.
- `POST /accounts/transfer`: Transfer money between two accounts.
- `GET /accounts/:id/transactions`: Get transaction history for an account.

### Loans
- `POST /loans`: Apply for a new loan.
- `GET /loans/:id`: Get loan details (includes interest calculation).
- `POST /loans/:id/repay`: Repay a loan.

## Setup & Installation

### Prerequisites
- [Go](https://golang.org/dl/) (version 1.16 or higher)
- [PostgreSQL](https://www.postgresql.org/download/)

### Installation

1.  **Clone the repository:**
    ```sh
    git clone <repository-url>
    cd banking-system
    ```

2.  **Configure Environment Variables:**
    Create a `.env` file in the root directory with your database credentials:
    ```env
    DB_HOST=localhost
    DB_USER=your_postgres_user
    DB_PASSWORD=your_postgres_password
    DB_NAME=banking_system
    DB_PORT=5432
    DB_SSLMODE=disable
    ```

3.  **Install Dependencies:**
    ```sh
    go mod download
    ```

4.  **Run the Application:**
    ```sh
    go run cmd/server/main.go
    ```
    The server will start on `http://localhost:8080`.

## Project Structure

- `cmd/server`: Entry point of the application.
- `internal/config`: Configuration logic.
- `internal/db`: Database connection and migration.
- `internal/handlers`: HTTP request handlers (controllers).
- `internal/models`: Database models (gorm structs).
- `internal/repositories`: Data access layer.
- `internal/services`: Business logic layer.
