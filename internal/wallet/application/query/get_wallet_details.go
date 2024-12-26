package query

import (
	"context"

	"github.com/Rustixir/wms/internal/wallet/domain"
	"github.com/Rustixir/wms/internal/wallet/port/output"
	"github.com/google/uuid"
)

type GetWalletDetailsQuery struct {
	WalletID uuid.UUID
}

type GetWalletDetailsHandler struct {
	walletRepo output.WalletRepo
}

func NewGetWalletDetailsHandler(walletRepo output.WalletRepo) *GetWalletDetailsHandler {
	return &GetWalletDetailsHandler{
		walletRepo: walletRepo,
	}
}

func (h *GetWalletDetailsHandler) Handle(
	ctx context.Context,
	qry GetWalletDetailsQuery,
) (domain.Wallet, error) {
	wallet, err := h.walletRepo.FindById(ctx, qry.WalletID)
	if err != nil {
		return domain.Wallet{}, err
	}
	return wallet, nil
}
