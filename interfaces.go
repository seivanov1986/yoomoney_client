package yoomoney_client

type Yoomoney interface {
	GenPayment(
		summ string,
		description string,
		email string,
		phone string,
	) string
}
