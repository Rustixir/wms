package main

import (
	"database/sql"
	errors2 "errors"
	"net/http"
	"strings"

	"github.com/Rustixir/wms/internal/wallet/adapter/input/rest"
	"github.com/Rustixir/wms/internal/wallet/adapter/output/persistence"
	"github.com/Rustixir/wms/internal/wallet/application/command"
	"github.com/Rustixir/wms/internal/wallet/application/query"
	"github.com/Rustixir/wms/internal/wallet/config"
	"github.com/Rustixir/wms/pkg/errors"
	"github.com/Rustixir/wms/pkg/localization"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := sql.Open("database/sql", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	echo := echo.New()
	group := echo.Group("/v1")
	echo.HTTPErrorHandler = ErrorHandler

	// initialize the repository
	repo := persistence.NewPgWalletRepo(db)

	// register the handlers
	rest.NewWalletHandler(
		command.NewCreateWalletHandler(repo),
		command.NewAddFundsHandler(repo),
		command.NewDeductFundsHandler(repo),
		command.NewBlockWalletHandler(repo),
		command.NewUnblockWalletHandler(repo),
		query.NewGetWalletDetailsHandler(repo),
		query.NewGetTransactionHistoryHandler(repo),
	).Register(group)
	echo.Logger.Fatal(echo.Start(":" + config.Object.Port))
}

func ErrorHandler(err error, c echo.Context) {
	val := c.Request().Header.Get("Accept-Language")
	var lang string = "en"
	if strings.Contains(val, "fa") {
		lang = "fa"
	}

	var keyErr errors.KeyError
	if errors2.As(err, &keyErr) {
		c.NoContent(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"error": localization.Get(lang, keyErr.Key),
	})
}
