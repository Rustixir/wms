package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID
	WalletID  uuid.UUID
	Amount    float64
	Type      TransactionType
	Status    TransactionStatus
	Timestamp time.Time
}

// --------------------- constants  ----------------------------

type TransactionType string

const (
	TransactionTypeCredit TransactionType = "credit"
	TransactionTypeDebit  TransactionType = "debit"
)

type TransactionStatus string

const (
	TransactionStatusPending   TransactionStatus = "pending"
	TransactionStatusCompleted TransactionStatus = "completed"
)
