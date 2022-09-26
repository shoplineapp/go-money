package money

import "errors"

var (
	// ErrCurrencyMismatch happens when two compared Money don't have the same currency.
	ErrCurrencyMismatch = errors.New("currencies don't match")
)

type Money struct {
	Cents          int64   `json:"cents" bson:"cents"`
	CurrencySymbol string  `json:"currency_symbol" bson:"currency_symbol"`
	CurrencyIso    string  `json:"currency_iso" bson:"currency_iso"`
	Label          string  `json:"label" bson:"label"`
	Dollars        float64 `json:"dollars" bson:"dollars"`

	currency *Currency
}

func New(cent int64, code string) *Money {
	return &Money{
		Cents:    cent,
		currency: NewCurrency(code),
	}
}

func (m *Money) loadConfig() {
	if m.currency == nil {
		m.currency = NewCurrency(m.CurrencyIso)
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

func (m *Money) Multiply(multiplicand int64) *Money {
	m.loadConfig()
	cent := m.Cents * multiplicand
	return New(cent, m.CurrencyIso)
}

func (m *Money) Divide(dividend int64) (*Money, error) {
	m.loadConfig()
	return nil, nil
}

func (m *Money) IsSameCurrency(om *Money) bool {
	return m.CurrencyIso == om.CurrencyIso
}

func (m *Money) assertSameCurrency(om *Money) error {
	if !m.IsSameCurrency(om) {
		return ErrCurrencyMismatch
	}

	return nil
}

func (m *Money) compare(om *Money) int {
	switch {
	case m.Cents > om.Cents:
		return 1
	case m.Cents < om.Cents:
		return -1
	}

	return 0
}

// Equals checks equality between two Money types.
func (m *Money) Equals(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) == 0, nil
}

// GreaterThan checks whether the value of Money is greater than the other.
func (m *Money) GreaterThan(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) == 1, nil
}

// GreaterThanOrEqual checks whether the value of Money is greater or equal than the other.
func (m *Money) GreaterThanOrEqual(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) >= 0, nil
}

// LessThan checks whether the value of Money is less than the other.
func (m *Money) LessThan(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) == -1, nil
}

// LessThanOrEqual checks whether the value of Money is less or equal than the other.
func (m *Money) LessThanOrEqual(om *Money) (bool, error) {
	if err := m.assertSameCurrency(om); err != nil {
		return false, err
	}

	return m.compare(om) <= 0, nil
}

func (m *Money) GetLabel() string {
	// TODO format label
	return m.Label
}
