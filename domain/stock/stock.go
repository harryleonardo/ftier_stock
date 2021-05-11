package stock

import "github.com/labstack/echo"

type Usecase interface {
	StockProcessing(ctx echo.Context) (interface{}, error)
}
