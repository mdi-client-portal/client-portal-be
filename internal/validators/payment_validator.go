package validators

type PaymentClientValidator struct {
	ClientId string `json:"client_id" validate:"required"`
}
