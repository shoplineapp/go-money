package money

import (
	gomoney "github.com/Rhymond/go-money"
	"math"
	"math/big"
)

type Money struct {
	Cents          int64   `json:"cents" bson:"cents"`
	CurrencySymbol string  `json:"currency_symbol" bson:"currency_symbol"`
	CurrencyIso    string  `json:"currency_iso" bson:"currency_iso"`
	Label          string  `json:"label" bson:"label"`
	Dollars        float64 `json:"dollars" bson:"dollars"`

	money        *gomoney.Money
	roundingMode big.RoundingMode
}

func (m *Money) initMoney() {
	if m.money == nil {
		m.money = gomoney.New(m.Cents, m.CurrencyIso)
	}
}

func (m *Money) SetRoundMode(rm big.RoundingMode) {
	m.roundingMode = rm
}

// Equals checks equality between two Money types.
func (m *Money) Equals(om *Money) (bool, error) {
	m.initMoney()
	om.initMoney()
	return m.money.Equals(om.money)
}

// GreaterThan checks whether the value of Money is greater than the other.
func (m *Money) GreaterThan(om *Money) (bool, error) {
	m.initMoney()
	om.initMoney()
	return m.money.GreaterThan(om.money)
}

// GreaterThanOrEqual checks whether the value of Money is greater or equal than the other.
func (m *Money) GreaterThanOrEqual(om *Money) (bool, error) {
	m.initMoney()
	om.initMoney()
	return m.money.GreaterThanOrEqual(om.money)
}

// LessThan checks whether the value of Money is less than the other.
func (m *Money) LessThan(om *Money) (bool, error) {
	m.initMoney()
	om.initMoney()
	return m.money.LessThan(om.money)
}

// LessThanOrEqual checks whether the value of Money is less or equal than the other.
func (m *Money) LessThanOrEqual(om *Money) (bool, error) {
	m.initMoney()
	om.initMoney()
	return m.money.LessThanOrEqual(om.money)
}

// IsZero returns boolean of whether the value of Money is equals to zero.
func (m *Money) IsZero() bool {
	m.initMoney()
	return m.money.IsZero()
}

// IsPositive returns boolean of whether the value of Money is positive.
func (m *Money) IsPositive() bool {
	m.initMoney()
	return m.money.IsPositive()
}

// IsNegative returns boolean of whether the value of Money is negative.
func (m *Money) IsNegative() bool {
	m.initMoney()
	return m.money.IsNegative()
}

// Absolute returns new Money struct from given Money using absolute monetary value.
func (m *Money) Absolute() *Money {
	m.initMoney()
	return m.Absolute()
}

// Negative returns new Money struct from given Money using negative monetary value.
func (m *Money) Negative() *Money {
	m.initMoney()
	return m.Negative()
}

// Add returns new Money struct with value representing sum of Self and Other Money.
func (m *Money) Add(om *Money) (*Money, error) {
	m.initMoney()
	om.initMoney()

	nm, err := m.money.Add(om.money)
	if err != nil {
		return nil, err
	}

	return &Money{
		Cents:          nm.Amount(),
		Dollars:        nm.AsMajorUnits(),
		CurrencyIso:    m.CurrencyIso,
		CurrencySymbol: m.CurrencySymbol,
	}, nil
}

// Subtract returns new Money struct with value representing difference of Self and Other Money.
func (m *Money) Subtract(om *Money) (*Money, error) {
	m.initMoney()
	om.initMoney()

	nm, err := m.money.Subtract(om.money)
	if err != nil {
		return nil, err
	}

	return &Money{
		Cents:          nm.Amount(),
		Dollars:        nm.AsMajorUnits(),
		CurrencyIso:    m.CurrencyIso,
		CurrencySymbol: m.CurrencySymbol,
	}, nil
}

// Multiply returns new Money struct with value representing Self multiplied value by multiplier.
func (m *Money) Multiply(mul float64) *Money {
	m.initMoney()

	cents := m.money.Amount()
	newCents := float64(cents) * mul
	// TODO support rounding mode
	round := math.RoundToEven(newCents)

	nm := gomoney.New(int64(round), m.CurrencyIso)

	return &Money{
		Cents:          nm.Amount(),
		Dollars:        nm.AsMajorUnits(),
		CurrencyIso:    m.CurrencyIso,
		CurrencySymbol: m.CurrencySymbol,
	}
}
