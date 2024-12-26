package query

import (
	"context"
	"log/slog"

	"github.com/Rustixir/wms/internal/wallet/domain"
	"github.com/Rustixir/wms/internal/wallet/port/output"
	"github.com/Rustixir/wms/pkg/errors"
	"github.com/google/uuid"
)

type GetTransactionHistoryQuery struct {
	WalletID uuid.UUID
	pageSize int
	pageNum  int
}

type GetTransactionHistoryHandler struct {
	walletRepo output.WalletRepo
}

func NewGetTransactionHistoryHandler(walletRepo output.WalletRepo) *GetTransactionHistoryHandler {
	return &GetTransactionHistoryHandler{
		walletRepo: walletRepo,
	}
}

func (h *GetTransactionHistoryHandler) Handle(
	ctx context.Context,
	qry GetTransactionHistoryQuery,
) ([]domain.Transaction, error) {
	limit := qry.pageSize
	offset := (qry.pageNum - 1) * qry.pageSize
	list, err := h.walletRepo.GetTransactionHistory(ctx, qry.WalletID, limit, offset)
	if err != nil {
		slog.Error("GetTransactionHistoryHandler.handle", "error", err)
		return nil, errors.ErrWalletNotfound
	}
	return list, nil
}
