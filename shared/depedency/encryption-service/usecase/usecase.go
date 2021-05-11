package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ftier-stock/shared/config"
	"github.com/ftier-stock/shared/depedency/encryption-service"
	"github.com/ftier-stock/shared/depedency/encryption-service/vo"
	"github.com/parnurzeal/gorequest"
)

type usecase struct {
	// - inject some depedency if needed;
	config config.ImmutableConfig
}

func NewEncrypt(config config.ImmutableConfig) encryption.Usecase {
	return &usecase{
		config: config,
	}
}

func (u usecase) EncryptMessage(dto *vo.EncryptRequest) (*vo.EncryptResponse, error) {
	payload, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s", u.config.GetEncryptionSvcURL())
	res, body, requestErr := gorequest.
		New().
		Post(url).
		Set("Content-Type", "application/json").
		Send(string(payload)).
		End()

	if requestErr != nil {
		return nil, fmt.Errorf("Got error while trying to call depedency; Error message :%s", requestErr)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Got error response : %v", res)
	}

	var encryptResponse vo.EncryptResponse
	err = json.Unmarshal([]byte(body), &encryptResponse)
	if err != nil {
		return nil, fmt.Errorf("Failed to UnMarshal data due to : ", err)
	}

	return &encryptResponse, nil
}
