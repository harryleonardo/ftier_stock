package usecase

import (
	"fmt"
	"net/http"

	"github.com/ftier-stock/shared/config"
	"github.com/ftier-stock/shared/depedency/alphavantage"
	"github.com/parnurzeal/gorequest"
)

type usecase struct {
	// - inject some depedency if needed;
	config config.ImmutableConfig
}

func NewAlphavantage(config config.ImmutableConfig) alphavantage.Usecase {
	return &usecase{config: config}
}

// - implement function that declare in usecase;
func (u usecase) GetSymbol(stockParam string) (interface{}, error) {
	url := fmt.Sprintf("%s/query?function=TIME_SERIES_INTRADAY&interval=1min&symbol=%s&apikey=%s", u.config.GetAlphavantageURL(), stockParam, u.config.GetAlphavantageKey())
	res, body, err := gorequest.
		New().
		Get(url).
		End()

	if err != nil {
		return nil, fmt.Errorf("Got error while trying to call depedency; Error message :%s", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Got error response : %v", res)
	}

	return body, nil
}
