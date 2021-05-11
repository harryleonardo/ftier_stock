package vo

type (
	EncryptRequest struct {
		Message interface{} `json:"message"`
	}

	EncryptResponse struct {
		Message interface{} `json:"encrypted_message"`
	}
)
