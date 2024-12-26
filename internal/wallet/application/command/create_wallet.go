package command

import (
	"context"
	"time"

	"github.com/Rustixir/wms/internal/wallet/domain"
	"github.com/Rustixir/wms/internal/wallet/port/output"
	"github.com/google/uuid"
)

type CreateWalletCommand struct {
	OwnerID  string
	Currency string
}

type CreateWalletHandler struct {
	walletRepo output.WalletRepo
}

func NewCreateWalletHandler(walletRepo output.WalletRepo) *CreateWalletHandler {
	return &CreateWalletHandler{
		walletRepo: walletRepo,
	}
}

func (h *CreateWalletHandler) Handle(ctx context.Context, command CreateWalletCommand) (err error) {
	return h.walletRepo.CreateWallet(ctx, domain.Wallet{
		ID:        uuid.New(),
		OwnerID:   command.OwnerID,
		Currency:  command.Currency,
		Balance:   0,
		Status:    domain.StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}
