package types

// Amount of money in minimum units (diram, kopeyka, cent...)
type Money int64

// Code of currency
type Currency string

// Codes of currency
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// Number of card
type PAN string

// Info about payment card
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
	MinBalance Money
}

type PaymentSource struct {
	Type    string
	Number  string
	Balance types.Money
}
