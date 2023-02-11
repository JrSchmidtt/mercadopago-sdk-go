package service

import (
	"encoding/json"

	"github.com/joelschutz/mercadopago-sdk-go"
	model "github.com/joelschutz/mercadopago-sdk-go/models"
	"github.com/joelschutz/mercadopago-sdk-go/request"
)

// CreateCardToken é o método responsável por criar um token de cartão no MercadoPago.
// Ou seja é o método que você deve usar para criar um token de cartão para ser usado em pagamentos futuros.
func CreateCardToken(cardTokenRequest model.CardTokenRequest, mercadoPagoAccessToken ...string) (*model.CardTokenResponse, *model.ErrorResponse, error) {
	params := request.Params{
		Method:  "POST",
		Headers: map[string]interface{}{"Authorization": "Bearer " + mercadopago.GetAccessToken(mercadoPagoAccessToken...)},
		URL    : mercadopago.BASEURL + "/v1/card_tokens",
	}

	response, err := request.New(params)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode > 300 {
		resp, err := mercadopago.ParseError(response.RawBody)
		return nil, resp, err
	}

	var cardTokenResponse model.CardTokenResponse
	err = json.Unmarshal(response.RawBody, &cardTokenResponse)
	return &cardTokenResponse, nil, err
}