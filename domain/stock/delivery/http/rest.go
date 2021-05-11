package http

import (
	"net/http"

	"github.com/ftier-stock/domain/stock"
	"github.com/labstack/echo"
)

type stockHandler struct {
	usecase stock.Usecase
}

func StockHandler(e *echo.Echo, usecase stock.Usecase) {
	handler := stockHandler{
		usecase: usecase,
	}

	e.GET("api.finantier.test/symbol", handler.GetStockSymbol)
}

func (handler stockHandler) GetStockSymbol(e echo.Context) error {
	res, err := handler.usecase.StockProcessing(e)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, res)
}
