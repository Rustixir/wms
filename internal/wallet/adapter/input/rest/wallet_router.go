package rest

import (
	"github.com/Rustixir/wms/internal/wallet/application/command"
	"github.com/Rustixir/wms/internal/wallet/application/query"
	"github.com/labstack/echo/v4"
)

type WalletHandler struct {
	CreateWalletHandler          *command.CreateWalletHandler
	AddFundsHandler              *command.AddFundsHandler
	DeductFundsHandler           *command.DeductFundsHandler
	BlockWalletHandler           *command.BlockWalletHandler
	UnblockWalletHandler         *command.UnblockWalletHandler
	GetWalletDetailsHandler      *query.GetWalletDetailsHandler
	GetTransactionHistoryHandler *query.GetTransactionHistoryHandler
}

func NewWalletHandler(
	createWalletHandler *command.CreateWalletHandler,
	addFundsHandler *command.AddFundsHandler,
	deductFundsHandler *command.DeductFundsHandler,
	blockWalletHandler *command.BlockWalletHandler,
	unblockWalletHandler *command.UnblockWalletHandler,
	getWalletDetailsHandler *query.GetWalletDetailsHandler,
	getTransactionHistoryHandler *query.GetTransactionHistoryHandler,
) *WalletHandler {
	return &WalletHandler{
		CreateWalletHandler:          createWalletHandler,
		AddFundsHandler:              addFundsHandler,
		DeductFundsHandler:           deductFundsHandler,
		BlockWalletHandler:           blockWalletHandler,
		UnblockWalletHandler:         unblockWalletHandler,
		GetWalletDetailsHandler:      getWalletDetailsHandler,
		GetTransactionHistoryHandler: getTransactionHistoryHandler,
	}
}

func (h *WalletHandler) Register(g *echo.Group) {
	g.POST("/wallet", h.CreateWallet)
	g.POST("/wallet/:walletID/funds", h.AddFunds)
	g.POST("/wallet/:walletID/deduct", h.DeductFunds)
	g.POST("/wallet/:walletID/block", h.BlockWallet)
	g.POST("/wallet/:walletID/unblock", h.UnblockWallet)
	g.GET("/wallet/:walletID", h.GetWalletDetails)
	g.GET("/wallet/:walletID/transactions", h.GetTransactionHistory)
}
