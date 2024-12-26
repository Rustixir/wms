package command

import (
	"context"
	"log/slog"

	"github.com/Rustixir/wms/internal/wallet/port/output"
	"github.com/google/uuid"
)

type UnblockWalletCommand struct {
	WalletID uuid.UUID
}

type UnblockWalletHandler struct {
	walletRepo output.WalletRepo
}

func NewUnblockWalletHandler(walletRepo output.WalletRepo) *UnblockWalletHandler {
	return &UnblockWalletHandler{
		walletRepo: walletRepo,
	}
}

func (h *UnblockWalletHandler) Handle(ctx context.Context, command UnblockWalletCommand) (err error) {
	wallet, err := h.walletRepo.FindById(ctx, command.WalletID)
	if err != nil {
		slog.Error("UnblockWalletHandler.Handle", "error", err)
		return err
	}
	wallet.UnblockWallet()
	err = h.walletRepo.UpdateWallet(ctx, wallet)
	if err != nil {
		slog.Error("UnblockWalletHandler.Handle", "error", err)
		return err
	}
	return nil
}
