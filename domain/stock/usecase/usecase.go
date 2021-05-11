package usecase

import (
	"fmt"
	"net/http"

	"github.com/ftier-stock/domain/stock"
	SharedContext "github.com/ftier-stock/shared/context"
	"github.com/ftier-stock/shared/depedency/alphavantage"
	"github.com/ftier-stock/shared/depedency/encryption-service"
	EncryptionVO "github.com/ftier-stock/shared/depedency/encryption-service/vo"
	StockVO "github.com/ftier-stock/shared/vo"
	"github.com/labstack/echo"
)

type usecase struct {
	// - inject some depedency that will be used on this function;
	alphavantageDep  alphavantage.Usecase
	encryptionSvcDep encryption.Usecase
}

func NewStockUsecase(alphavantageDep alphavantage.Usecase, encryptionSvcDep encryption.Usecase) stock.Usecase {
	return &usecase{
		alphavantageDep:  alphavantageDep,
		encryptionSvcDep: encryptionSvcDep,
	}
}

func (u usecase) StockProcessing(ctx echo.Context) (interface{}, error) {
	ac := ctx.(*SharedContext.ApplicationContext)

	// - get stock name from parameter;
	stockName := ac.QueryParam("STOCK")

	// - pass the query param into 3rd party depedency (finantier);
	alphData, err := u.alphavantageDep.GetSymbol(stockName)
	if err != nil {
		return nil, err
	}

	// - logger for data that receive from alphavantage 3rd party API.
	fmt.Println("-------------------------------------------------------------------------------")
	fmt.Println("Stock Data : ", alphData)

	// - build payload that want to send into encryption service
	encryptionPayload := EncryptionVO.EncryptRequest{
		Message: alphData,
	}

	// - send the raw response data into 3rd party depedency (encryption service);
	encryptedMessage, err := u.encryptionSvcDep.EncryptMessage(&encryptionPayload)
	if err != nil {
		fmt.Println("Got Error from Encryption Service: ", err)
		return nil, ac.JSON(http.StatusInternalServerError, err)
	}

	return StockVO.StockResponse{
		EncryptedMessage: encryptedMessage.Message,
	}, nil
}
