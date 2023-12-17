package yoomoney_client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (y *yoomoney) GenPayment(
	summ string,
	description string,
	email string,
	phone string,
	returnUrl string,
) string {
	query := YoomoneyQuery{}
	query.Amount.Currency = "RUB"
	query.Amount.Value = summ
	query.Capture = true
	query.Confirmation.ReturnUrl = returnUrl
	query.Confirmation.Type = "redirect"
	query.Description = description
	query.Receipt.Customer.Email = email
	query.Receipt.Customer.Phone = phone
	query.Receipt.Id = uuid.Must(uuid.NewRandom()).String()
	query.Receipt.Type = "payment"
	query.Receipt.Status = "succeeded"
	query.Receipt.Items = []Item{
		{
			Description: description,
			Quantity:    "1.000",
			Amount: Amount{
				Value:    summ,
				Currency: "RUB",
			},
			VatCode:        "1",
			PaymentSubject: "service",
		},
	}

	idemp := strconv.FormatInt(time.Now().Unix(), 10)

	y.Curl.SetUrl(baseUrl + "/payments")
	y.Curl.SetMethod("POST")
	y.Curl.SetHeader(map[string][]string{
		"Idempotence-Key": {idemp},
		"Content-Type":    {"application/json"},
	})

	queryStr, _ := json.Marshal(query)

	y.Curl.SetPostData(string(queryStr))

	y.Curl.Exec()
	return string(y.Curl.GetResponse())
}
