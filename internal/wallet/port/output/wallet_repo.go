package output

import (
	"context"

	"github.com/Rustixir/wms/internal/wallet/domain"
	"github.com/google/uuid"
)

type WalletRepo interface {
	CreateWallet(ctx context.Context, wallet domain.Wallet) error
	UpdateWallet(ctx context.Context, wallet domain.Wallet) error
	FindById(ctx context.Context, walletID uuid.UUID) (domain.Wallet, error)
	Fetch(ctx context.Context, ownerID string) ([]domain.Wallet, error)
	CreateTransaction(ctx context.Context, transaction domain.Transaction) error
	GetTransactionHistory(ctx context.Context, WalletID uuid.UUID, limit int, offset int) ([]domain.Transaction, error)
}
