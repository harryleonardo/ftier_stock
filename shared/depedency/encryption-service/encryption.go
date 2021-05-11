package encryption

import "github.com/ftier-stock/shared/depedency/encryption-service/vo"

type Usecase interface {
	EncryptMessage(dto *vo.EncryptRequest) (*vo.EncryptResponse, error)
}
