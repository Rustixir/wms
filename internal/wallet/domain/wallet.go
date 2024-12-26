package domain

import (
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	ID        uuid.UUID
	OwnerID   string
	Currency  string
	Balance   float64
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

// AddFunds increases the wallet balance.
func (w *Wallet) AddFunds(amount float64) {
	w.Balance += amount
	w.UpdatedAt = time.Now()
}

// DeductFunds decreases the wallet balance.
func (w *Wallet) DeductFunds(amount float64) {
	w.Balance -= amount
	w.UpdatedAt = time.Now()
}

// BlockWallet blocks the wallet.
func (w *Wallet) BlockWallet() {
	w.Status = StatusBlocked
	w.UpdatedAt = time.Now()
}

// UnblockWallet unblocks the wallet.
func (w *Wallet) UnblockWallet() {
	w.Status = StatusActive
	w.UpdatedAt = time.Now()
}

// HaveSufficientFunds returns true if the wallet has sufficient funds.
func (w *Wallet) HaveSufficientFunds(amount float64) bool {
	return w.Balance >= amount
}

// IsBlocked returns true if the wallet is active.
func (w *Wallet) IsBlocked() bool {
	return w.Status == StatusBlocked
}

// --------------------- constants  ----------------------------

type Status string

const (
	StatusActive  Status = "active"
	StatusBlocked Status = "blocked"
)
