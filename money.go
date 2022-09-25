package money

type Money struct {
	Cents          int64   `json:"cents" bson:"cents"`
	CurrencySymbol string  `json:"currency_symbol" bson:"currency_symbol"`
	CurrencyIso    string  `json:"currency_iso" bson:"currency_iso"`
	Label          string  `json:"label" bson:"label"`
	Dollars        float64 `json:"dollars" bson:"dollars"`

	currency *Currency
}

func New(amount int64, code string) *Money {
	return &Money{
		Cents:    amount,
		currency: newCurrency(code),
	}
}

func (m *Money) loadConfig() {
	if m.currency == nil {
		m.currency = newCurrency(m.CurrencyIso)
	}
}

func (m *Money) Add(addend Money) (*Money, error) {
	m.loadConfig()
	return nil, nil
}

func (m *Money) Subtract(subtracted Money) (*Money, error) {
	m.loadConfig()
	return nil, nil
}

func (m *Money) Multiply(multiplicand int64) (*Money, error) {
	m.loadConfig()
	return nil, nil
}

func (m *Money) Divide(dividend int64) (*Money, error) {
	m.loadConfig()
	return nil, nil
}
