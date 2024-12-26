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

type DeductFundsCommand struct {
	WalletID uuid.UUID
	Amount   float64
}

type DeductFundsHandler struct {
	walletRepo output.WalletRepo
}

func NewDeductFundsHandler(walletRepo output.WalletRepo) *DeductFundsHandler {
	return &DeductFundsHandler{
		walletRepo: walletRepo,
	}
}

func (h *DeductFundsHandler) Handle(ctx context.Context, command DeductFundsCommand) (err error) {
	wallet, err := h.walletRepo.FindById(ctx, command.WalletID)
	if err != nil {
		slog.Error("DeductFundsHandler.Handle", "error", err)
		return errors.ErrWalletNotfound
	}

	if wallet.IsBlocked() {
		return errors.ErrBlocked
	}

	if !wallet.HaveSufficientFunds(command.Amount) {
		return errors.ErrInsufficientFunds
	}

	wallet.AddFunds(command.Amount)
	if err = h.walletRepo.UpdateWallet(ctx, wallet); err != nil {
		slog.Error("DeductFundsHandler.Handle", "error", err)
		return errors.ErrInternalServerError
	}

	err = h.walletRepo.CreateTransaction(ctx, domain.Transaction{
		ID:        uuid.New(),
		WalletID:  wallet.ID,
		Amount:    command.Amount,
		Type:      domain.TransactionTypeDebit,
		Status:    domain.TransactionStatusCompleted,
		Timestamp: time.Now(),
	})
	if err != nil {
		slog.Error("DeductFundsHandler.Handle", "error", err)
		return errors.ErrInternalServerError
	}
	return nil
}
