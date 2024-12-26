package rest

import (
	"net/http"

	"github.com/Rustixir/wms/internal/wallet/application/command"
	"github.com/Rustixir/wms/internal/wallet/application/query"
	"github.com/Rustixir/wms/pkg/errors"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (w *WalletHandler) CreateWallet(ctx echo.Context) error {
	var cmd command.CreateWalletCommand
	if err := ctx.Bind(&cmd); err != nil {
		return errors.ErrInvalidRequest
	}
	err := w.CreateWalletHandler.Handle(ctx.Request().Context(), cmd)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}

func (w *WalletHandler) AddFunds(ctx echo.Context) error {
	walletID, err := getWalletID(ctx)
	if err != nil {
		return err
	}
	cmd := command.AddFundsCommand{
		WalletID: walletID,
	}
	if err := ctx.Bind(&cmd); err != nil {
		return errors.ErrInvalidRequest
	}
	err = w.AddFundsHandler.Handle(ctx.Request().Context(), cmd)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}

func (w *WalletHandler) DeductFunds(ctx echo.Context) error {
	walletID, err := getWalletID(ctx)
	if err != nil {
		return err
	}
	cmd := command.DeductFundsCommand{
		WalletID: walletID,
	}
	if err := ctx.Bind(&cmd); err != nil {
		return errors.ErrInvalidRequest
	}
	err = w.DeductFundsHandler.Handle(ctx.Request().Context(), cmd)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}

func (w *WalletHandler) BlockWallet(ctx echo.Context) error {
	walletID, err := getWalletID(ctx)
	if err != nil {
		return err
	}
	cmd := command.BlockWalletCommand{
		WalletID: walletID,
	}
	if err := ctx.Bind(&cmd); err != nil {
		return errors.ErrInvalidRequest
	}
	err = w.BlockWalletHandler.Handle(ctx.Request().Context(), cmd)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}

func (w *WalletHandler) UnblockWallet(ctx echo.Context) error {
	walletID, err := getWalletID(ctx)
	if err != nil {
		return err
	}
	cmd := command.UnblockWalletCommand{
		WalletID: walletID,
	}
	if err := ctx.Bind(&cmd); err != nil {
		return errors.ErrInvalidRequest
	}
	err = w.UnblockWalletHandler.Handle(ctx.Request().Context(), cmd)
	if err != nil {
		return err
	}
	return ctx.NoContent(http.StatusCreated)
}

func (w *WalletHandler) GetWalletDetails(ctx echo.Context) error {
	walletID, err := getWalletID(ctx)
	if err != nil {
		return err
	}
	qry := query.GetWalletDetailsQuery{
		WalletID: walletID,
	}
	if err := ctx.Bind(&qry); err != nil {
		return errors.ErrInvalidRequest
	}
	res, err := w.GetWalletDetailsHandler.Handle(ctx.Request().Context(), qry)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func (w *WalletHandler) GetTransactionHistory(ctx echo.Context) error {
	walletID, err := getWalletID(ctx)
	if err != nil {
		return err
	}
	qry := query.GetTransactionHistoryQuery{
		WalletID: walletID,
	}
	if err := ctx.Bind(&qry); err != nil {
		return errors.ErrInvalidRequest
	}
	res, err := w.GetTransactionHistoryHandler.Handle(ctx.Request().Context(), qry)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, res)
}

func getWalletID(ctx echo.Context) (uuid.UUID, error) {
	val := ctx.Param("walletID")
	if val == "" {
		return uuid.Nil, errors.ErrInvalidWalletID
	}
	uid, err := uuid.Parse(val)
	if err != nil {
		return uuid.Nil, errors.ErrInvalidWalletID
	}
	return uid, nil
}
