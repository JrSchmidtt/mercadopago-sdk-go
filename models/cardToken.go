package model

// Resposta da API para o token do cartão
// https://www.mercadopago.com.br/developers/pt/reference/card_tokens/_card_tokens/post/
type CardTokenResponse struct {
	ID                 string     `json:"id,omitempty"`
	FirstSixDigits     string     `json:"first_six_digits,omitempty"`
	ExpirationMonth    int64      `json:"expiration_month,omitempty"`
	ExpirationYear     int64      `json:"expiration_year,omitempty"`
	LastFourDigits     string     `json:"last_four_digits,omitempty"`
	Cardholder         Cardholder `json:"cardholder,omitempty"`
	Status             string     `json:"status,omitempty"`
	DateCreated        string     `json:"date_created,omitempty"`
	DateLastUpdated    string     `json:"date_last_updated,omitempty"`
	DateDue            string     `json:"date_due,omitempty"`
	LuhnValidation     bool       `json:"luhn_validation,omitempty"`
	LiveMode           bool       `json:"live_mode,omitempty"`
	RequireEsc         bool       `json:"require_esc,omitempty"`
	CardNumberLength   int64      `json:"card_number_length,omitempty"`
	SecurityCodeLength int64      `json:"security_code_length,omitempty"`
}

// Requisição da API para o token do cartão
// https://www.mercadopago.com.br/developers/pt/reference/card_tokens/_card_tokens/post/
type CardTokenRequest struct {
	CardNumber      string     `json:"card_number,omitempty"`
	ExpirationMonth int64      `json:"expiration_month,omitempty"`
	ExpirationYear  int64      `json:"expiration_year,omitempty"`
	SecurityCode    string     `json:"security_code,omitempty"`
	Cardholder      Cardholder `json:"cardholder,omitempty"`
}