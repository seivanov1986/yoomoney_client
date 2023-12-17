package yoomoney_client

import (
	"github.com/seivanov1986/http_client"
)

const (
	baseUrl = "https://api.yookassa.ru/v3"
)

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type PaymentMethodData struct {
	Type string `json:"type,omitempty"`
}

type Confirmation struct {
	Type      string `json:"type"`
	ReturnUrl string `json:"return_url"`
}

type Customer struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Item struct {
	Description    string `json:"description"`
	Quantity       string `json:"quantity"`
	VatCode        string `json:"vat_code"`
	Amount         Amount `json:"amount"`
	PaymentSubject string `json:"payment_subject"`
}

type Receipt struct {
	Customer Customer `json:"customer"`
	Id       string   `json:"id"`
	Type     string   `json:"type"`
	Status   string   `json:"status"`
	Items    []Item   `json:"items"`
}

type YoomoneyQuery struct {
	Amount              Amount             `json:"amount"`
	PaymentMethodData   *PaymentMethodData `json:"payment_method_data,omitempty"`
	Capture             bool               `json:"capture"`
	Confirmation        Confirmation       `json:"confirmation"`
	Description         string             `json:"description"`
	ReceiptRegistration string             `json:"receipt_registration"`
	Receipt             Receipt            `json:"receipt"`
}

type yoomoney struct {
	Curl http_client.HttpClient
}

func NewYooMoneyClient(curl http_client.HttpClient) *yoomoney {
	return &yoomoney{
		Curl: curl,
	}
}
