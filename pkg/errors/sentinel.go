package errors

import "errors"

var (
	ErrWalletNotfound      = NewKeyError("wallet_not_found", errors.New("wallet not found"))
	ErrBlocked             = NewKeyError("wallet_blocked", errors.New("this wallet is blocked"))
	ErrInsufficientFunds   = NewKeyError("wallet_insufficient", errors.New("insufficient funds"))
	ErrInternalServerError = NewKeyError("internal_server_error", errors.New("internal server error"))
	ErrInvalidRequest      = NewKeyError("invalid_request", errors.New("invalid request"))
	ErrInvalidWalletID     = NewKeyError("invalid_wallet_id", errors.New("invalid wallet id"))
)
