package persistence

import (
	"context"
	"database/sql"

	"github.com/Rustixir/wms/internal/wallet/domain"
	"github.com/google/uuid"
)

func NewPgWalletRepo(db *sql.DB) *PgWalletRepo {
	return &PgWalletRepo{
		db: db,
	}
}

type PgWalletRepo struct {
	db *sql.DB
}

func (p *PgWalletRepo) CreateWallet(ctx context.Context, wallet domain.Wallet) error {
	query := `
		INSERT INTO wallets(
		   id, owner_id, currency, balance, status, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := p.db.ExecContext(ctx, query,
		wallet.ID,
		wallet.OwnerID,
		wallet.Currency,
		wallet.Balance,
		wallet.Status,
		wallet.CreatedAt,
		wallet.UpdatedAt,
	)
	return err
}

func (p *PgWalletRepo) UpdateWallet(ctx context.Context, wallet domain.Wallet) error {
	query := `
		UPDATE walllets 
		Set balance = $1, status = $2, updated_at = $3
		WHERE id = $4
	`
	_, err := p.db.ExecContext(ctx, query, wallet.Balance, wallet.Status, wallet.UpdatedAt, wallet.ID)
	return err
}

func (p *PgWalletRepo) FindById(ctx context.Context, walletID uuid.UUID) (wallet domain.Wallet, err error) {
	query := `
	    SELECT 
    		id, owner_id, currency, balance, status, created_at, updated_at 
		FROM wallets WHERE id = $1`
	err = p.db.QueryRowContext(ctx, query, walletID).Scan(
		&wallet.ID,
		&wallet.OwnerID,
		&wallet.Currency,
		&wallet.Balance,
		&wallet.Status,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (p *PgWalletRepo) Fetch(ctx context.Context, ownerID string) ([]domain.Wallet, error) {
	query := `
	SELECT 
    	id, owner_id, currency, balance, status, created_at, updated_at 
	FROM wallets WHERE owner_id = $1`
	rows, err := p.db.QueryContext(ctx, query, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var wallets []domain.Wallet
	for rows.Next() {
		var wallet domain.Wallet
		err := rows.Scan(
			&wallet.ID,
			&wallet.OwnerID,
			&wallet.Currency,
			&wallet.Balance,
			&wallet.Status,
			&wallet.CreatedAt,
			&wallet.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet)
	}
	return wallets, nil
}

func (p *PgWalletRepo) CreateTransaction(ctx context.Context, transaction domain.Transaction) error {
	query := `
		INSERT INTO transactions(
		   id, wallet_id, amount, type, status, timestamp
		) VALUES ($1, $2, $3, $4, $5, $6)
		`
	_, err := p.db.ExecContext(ctx, query,
		transaction.ID,
		transaction.WalletID,
		transaction.Amount,
		transaction.Type,
		transaction.Status,
		transaction.Timestamp,
	)
	return err
}

func (p *PgWalletRepo) GetTransactionHistory(ctx context.Context, walletID uuid.UUID, limit int, offset int) ([]domain.Transaction, error) {
	query := `
		SELECT 
    		id, wallet_id, amount, type, status, timestamp 
		FROM transactions WHERE wallet_id = $1 ORDER BY timestamp DESC LIMIT $2 OFFSET $3`
	rows, err := p.db.QueryContext(ctx, query, walletID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []domain.Transaction
	for rows.Next() {
		var transaction domain.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.WalletID,
			&transaction.Amount,
			&transaction.Type,
			&transaction.Status,
			&transaction.Timestamp,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
