package command

import (
	"context"
	"log/slog"
	"time"

	"github.com/Rustixir/wms/internal/wallet/domain"
	"github.com/Rustixir/wms/internal/wallet/port/output"
	"github.com/Rustixir/wms/pkg/errors"
	"github.com/google/uuid"
)

type AddFundsCommand struct {
	WalletID uuid.UUID
	Amount   float64
}

type AddFundsHandler struct {
	walletRepo output.WalletRepo
}

func NewAddFundsHandler(walletRepo output.WalletRepo) *AddFundsHandler {
	return &AddFundsHandler{
		walletRepo: walletRepo,
	}
}

func (h *AddFundsHandler) Handle(ctx context.Context, command AddFundsCommand) (err error) {
	var wallet domain.Wallet

	wallet, err = h.walletRepo.FindById(ctx, command.WalletID)
	if err != nil {
		slog.Error("AddFundsHandler.Handle", "error", err)
		return errors.ErrWalletNotfound
	}

	if wallet.IsBlocked() {
		return errors.ErrBlocked
	}

	wallet.AddFunds(command.Amount)
	err = h.walletRepo.UpdateWallet(ctx, wallet)
	if err != nil {
		slog.Error("AddFundsHandler.Handle", "error", err)
		return errors.ErrWalletNotfound
	}

	err = h.walletRepo.CreateTransaction(ctx, domain.Transaction{
		ID:        uuid.New(),
		WalletID:  wallet.ID,
		Amount:    command.Amount,
		Type:      domain.TransactionTypeCredit,
		Status:    domain.TransactionStatusCompleted,
		Timestamp: time.Now(),
	})
	if err != nil {
		slog.Error("AddFundsHandler.Handle", "error", err)
		return errors.ErrInternalServerError
	}
	return nil
}
