package input

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type WalletService interface {
	CreateWallet(ctx context.Context, ownerID string, currency string) (uuid.UUID, error)
	AddFunds(ctx context.Context, walletID uuid.UUID, amount float64) error
	DeductFunds(ctx context.Context, walletID uuid.UUID, amount float64) error
	BlockWallet(ctx context.Context, walletID uuid.UUID) error
	UnblockWallet(ctx context.Context, walletID uuid.UUID) error
	GetWalletDetails(ctx context.Context, walletID uuid.UUID) (WalletDetails, error)
	GetTransactionHistory(ctx context.Context, walletID uuid.UUID, pageSize int, pageNum int) ([]TransactionDetails, error)
}

type WalletDetails struct {
	WalletID  uuid.UUID
	Balance   float64
	Currency  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TransactionDetails struct {
	TransactionID uuid.UUID
	WalletID      uuid.UUID
	Amount        float64
	Type          string
	Status        string
	Timestamp     time.Time
}
