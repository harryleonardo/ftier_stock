package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	SharedConfig "github.com/ftier-stock/shared/config"
	SharedContext "github.com/ftier-stock/shared/context"
	SharedAlphavantage "github.com/ftier-stock/shared/depedency/alphavantage/usecase"
	SharedEncryptionSvc "github.com/ftier-stock/shared/depedency/encryption-service/usecase"

	stockHandler "github.com/ftier-stock/domain/stock/delivery/http"
	stockUsecase "github.com/ftier-stock/domain/stock/usecase"
)

func main() {
	// - initialize echo labstack as a framework that i'm using;
	e := echo.New()

	// - initiate config;
	conf := SharedConfig.GetDefaultImmutableConfig()

	// - CORS
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// - initialize customize context;
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &SharedContext.ApplicationContext{
				Context: c,
				Config:  conf,
			}
			return h(ac)
		}
	})

	alphavantageDep := SharedAlphavantage.NewAlphavantage(conf)
	encryptionSvcDep := SharedEncryptionSvc.NewEncrypt(conf)

	stockUcase := stockUsecase.NewStockUsecase(alphavantageDep, encryptionSvcDep)

	stockHandler.StockHandler(e, stockUcase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.GetPort())))
}
