package command

import (
	"context"
	"log/slog"

	"github.com/Rustixir/wms/internal/wallet/port/output"
	"github.com/Rustixir/wms/pkg/errors"
	"github.com/google/uuid"
)

type BlockWalletCommand struct {
	WalletID uuid.UUID
}

type BlockWalletHandler struct {
	walletRepo output.WalletRepo
}

func NewBlockWalletHandler(walletRepo output.WalletRepo) *BlockWalletHandler {
	return &BlockWalletHandler{
		walletRepo: walletRepo,
	}
}

func (h *BlockWalletHandler) Handle(ctx context.Context, command BlockWalletCommand) (err error) {
	waller, err := h.walletRepo.FindById(ctx, command.WalletID)
	if err != nil {
		slog.Error("BlockWalletHandler.Handle", "error", err)
		return errors.ErrWalletNotfound
	}
	waller.BlockWallet()
	err = h.walletRepo.UpdateWallet(ctx, waller)
	if err != nil {
		slog.Error("BlockWalletHandler.Handle", "error", err)
		return errors.ErrInternalServerError
	}
	return nil
}
