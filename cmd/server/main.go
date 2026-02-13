package main

import (
	"banking-system/internal/config"
	"banking-system/internal/db"
	"banking-system/internal/handlers"
	"banking-system/internal/repositories"
	"banking-system/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	// Capture DB connection
	dbConn := db.Connect(cfg.DBDsn)

	// Ensure schema is up to date without dropping existing data.
	db.AutoMigrate(dbConn)

	// Repositories
	accountRepo := repositories.NewAccountRepository(dbConn)
	transactionRepo := repositories.NewTransactionRepository(dbConn)
	loanRepo := repositories.NewLoanRepository(dbConn)
	loanPaymentRepo := repositories.NewLoanPaymentRepository(dbConn)
	bankRepo := repositories.NewBankRepository(dbConn)
	branchRepo := repositories.NewBranchRepository(dbConn)
	customerRepo := repositories.NewCustomerRepository(dbConn)

	// Services
	accountService := services.NewAccountService(accountRepo, transactionRepo, customerRepo, branchRepo, bankRepo, dbConn)
	loanService := services.NewLoanService(loanRepo, loanPaymentRepo, customerRepo, bankRepo, branchRepo, dbConn)
	bankService := services.NewBankService(bankRepo)
	branchService := services.NewBranchService(branchRepo)
	customerService := services.NewCustomerService(customerRepo)

	// Handlers
	accountHandler := handlers.NewAccountHandler(accountService)
	loanHandler := handlers.NewLoanHandler(loanService)
	bankHandler := handlers.NewBankHandler(bankService)
	branchHandler := handlers.NewBranchHandler(branchService)
	customerHandler := handlers.NewCustomerHandler(customerService)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	// Bank routes
	r.POST("/banks", bankHandler.CreateBank)
	r.GET("/banks", bankHandler.GetBanks)

	// Branch routes
	r.POST("/branches", branchHandler.CreateBranch)
	r.GET("/branches", branchHandler.GetBranches)

	// Customer routes
	r.POST("/customers", customerHandler.CreateCustomer)
	r.GET("/customers/:id", customerHandler.GetCustomer)

	// Account routes
	r.POST("/accounts", accountHandler.OpenAccount)
	r.POST("/accounts/:id/deposit", accountHandler.Deposit)
	r.POST("/accounts/:id/withdraw", accountHandler.Withdraw)
	r.GET("/accounts/:id", accountHandler.GetAccount)
	r.GET("/accounts/:id/transactions", accountHandler.GetTransactions)
	r.POST("/accounts/transfer", accountHandler.Transfer)

	// Loan routes
	r.POST("/loans", loanHandler.CreateLoan)
	r.POST("/loans/:id/repay", loanHandler.RepayLoan)
	r.GET("/loans/:id", loanHandler.GetLoanDetails)

	r.Run(":8080")
}
